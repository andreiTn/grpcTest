package main

import (
	"myGrpc/api"
	"google.golang.org/grpc"
	"context"
	"net"

	"image"
	"image/color"
	"log"

	"github.com/disintegration/imaging"
	"fmt"
	"os"
	_ "google.golang.org/grpc/encoding/gzip"
	"io"
	"strings"
	"time"
)

const Root = "web/"
const UploadPath = "tmpimg/"

type Server struct {
	fileName string
	fullPath string
	newPath string
	oldPath string
}

func (s *Server) GetPaths(ctx context.Context, status *api.UploadStatus) (*api.NewImagePath, error) {
	return &api.NewImagePath{
		OldPath: s.oldPath,
		NewPath: s.newPath,
	}, nil
}

func (s *Server) Upload(stream api.ImageProcessor_UploadServer) (error) {
	s.oldPath = UploadPath+"v3-" + s.fileName
	s.fullPath = Root + UploadPath+"v3-" + s.fileName

	// open output file
	fo, err := os.Create(s.fullPath)

	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	log.Println("Started stream")

	for {
		fmt.Println("data from stream rcv")
		data, err := stream.Recv()

		if err != nil && err != io.EOF {
			fmt.Println("Som error::: ",err)
		}

		if err == io.EOF {
			fmt.Println("Goto end")
			goto END
		}

		_, writeErr := fo.Write(data.Content)
		if writeErr != nil {
			fmt.Println("write err")
		}
	}

END:
	err = s.GrayScale(stream)
	if err != nil {
		log.Fatalf("failed to send status code", err)
		return nil
	}

	return nil
}

func (s *Server) SendFileName(ctx context.Context, filName *api.FileName) (*api.FileName, error) {
	s.fileName = filName.Name

	return filName, nil
}

func (s *Server ) GrayScale(stream api.ImageProcessor_UploadServer) (error) {
	src, err := imaging.Open(s.fullPath)
	panicOn(err)

	img, err := os.Open(s.fullPath)
	panicOn(err)

	err = s.readAndStreamFile(s.fullPath + s.newPath, stream)
	panicOn(err)
	
	config, _, err := image.DecodeConfig(img)
	panicOn(err)

	// Create a grayscale version of the image with higher contrast and sharpness.
	grayscaleImg := imaging.Grayscale(src)

	dst := imaging.New(config.Width, config.Height, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, grayscaleImg, image.Pt(0, 0))


	// Save the resulting image using JPEG format.
	s.newPath = "img/v2------out_example.jpg"
	err = imaging.Save(dst, "web/" +s.newPath)
	panicOn(err)

	time.Sleep(time.Second * 3)
	fmt.Println("also need to stream this")
	err = s.readAndStreamFile("web/" + s.newPath, stream)
	panicOn(err)

	return nil
}

func (s *Server) readAndStreamFile(path string, stream api.ImageProcessor_UploadServer) error {
	file, err := os.Open(path)
	filename := strings.Split(path, "/")

	panicOn(err)

	// Close file and check for error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 50 * 1024)

	for  {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error while copying from file to buf", err)
			return nil
		}
		fmt.Println("Stream file: ", filename[len(filename)-1])
		err = stream.Send(&api.FileStream{Content: buf[:n], FileName: filename[len(filename)-1]})
		panicOn(err)

		return nil
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8888")
	fmt.Println("start...")

	if err != nil {
		fmt.Println("asd %s", err)
	}

	grpcSrv := grpc.NewServer()
	api.RegisterImageProcessorServer(grpcSrv, &Server{})
	grpcSrv.Serve(lis)

}

func panicOn(err error)  {
	if err != nil {
		log.Fatalf("Err: %s", err)
	}
}