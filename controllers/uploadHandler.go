package controllers

import "net/http"

func UploadHandler(w http.ResponseWriter, r *http.Request) {

}
//
//import (
//	"log"
//	"time"
//	"net/http"
//	"fmt"
//	"myGrpc/flashes"
//	"myGrpc/api"
//	"myGrpc/forms"
//	"mime/multipart"
//	"context"
//	"google.golang.org/grpc"
//	"io"
//	"sync"
//	"os"
//)
//
//type Stats struct {
//	StartedAt  time.Time
//	FinishedAt time.Time
//}
//
//// UploadFile uploads a file to the server
//func FileHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Redirect(w, r, "/", http.StatusSeeOther)
//		return
//	}
//
//	file, handle, err := r.FormFile(forms.UploadInputName)
//
//	if err != nil {
//		fmt.Fprintf(w, "Handle file err: %v", err)
//		return
//	}
//
//	defer file.Close()
//
//	mimeType := handle.Header.Get("Content-Type")
//	fmt.Println("File type: ", mimeType)
//	switch mimeType {
//	case "image/png", "image/jpeg":
//		save(w, r, file, handle)
//	default:
//		flash := []byte("The format file is not valid.")
//		flashes.SetFlash(w, "message", flash)
//
//		http.Redirect(w, r, "/", 301)
//	}
//}
//
//func save(w http.ResponseWriter, r *http.Request, file multipart.File, handle *multipart.FileHeader) {
//	var opts []grpc.DialOption
//	opts = append(opts, grpc.WithInsecure())
//
//	connection, err := grpc.Dial("localhost:8888", opts...)
//	defer connection.Close()
//
//	if err != nil {
//		log.Fatalf("Could not establish connection: %s", err)
//	}
//
//	client := api.NewImageProcessorClient(connection)
//
//	client.SendFileName(context.Background(), &api.FileName{
//		Name: handle.Filename,
//	})
//
//	os.Create("web/tmpimg/server{")
//
//	upload(context.Background() , file, client)
//
//	if err != nil {
//		log.Fatalf("Process err: %s", err)
//	}
//
//	flash := []byte("File saved")
//	flashes.SetFlash(w, "message", flash)
//
//	http.Redirect(w, r, "/show-image", 301)
//}
//
//func upload(ctx context.Context, file multipart.File, client api.ImageProcessorClient) (stats Stats, err error) {
//	var (
//		writing = true
//		buf     []byte
//		n       int
//	)
//
//	if err != nil {
//		log.Fatalf("failed to open file %s", err)
//		return
//	}
//	defer file.Close()
//
//	stream, err := client.Upload(ctx)
//
//	// open output file
//	fo, err := os.Create(file)
//
//	if err != nil {
//		panic(err)
//	}
//
//	// close fo on exit and check for its returned error
//	defer func() {
//		if err := fo.Close(); err != nil {
//			panic(err)
//		}
//	}()
//
//	log.Println("Started stream")
//
//	for {
//		fmt.Println("data from stream rcv")
//		data, err := stream.Recv()
//
//		if err != nil && err != io.EOF {
//			fmt.Println("Som error::: ",err)
//		}
//
//		if err == io.EOF {
//			fmt.Println("Goto end")
//			goto END
//		}
//
//		_, writeErr := fo.Write(data.Content)
//		if writeErr != nil {
//			fmt.Println("write err")
//		}
//	}
//
//	if err != nil {
//		log.Fatalf("failed to create upload stream for file %s", err)
//		return
//	}
//
//	stats.StartedAt = time.Now()
//	buf = make([]byte, 50 * 1024)
//
//	var wg sync.WaitGroup
//
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		for writing {
//			n, err = file.Read(buf)
//			if err != nil {
//				if err == io.EOF {
//					writing = false
//					err = nil
//					continue
//				}
//
//				log.Fatalf("errored while copying from file to buf", err)
//				return
//			}
//
//			err = stream.Send(&api.Chunk{
//				Content: buf[:n],
//			})
//			if err != nil {
//				log.Fatalf("failed to send chunk via stream", err)
//				return
//			}
//		}
//		stream.CloseSend()
//	}()
//
//	var filecreated bool
//	go func() {
//		defer wg.Done()
//
//		for {
//			fmt.Println("data from stream rcv")
//			data, err := stream.Recv()
//
//			if err != nil && err != io.EOF {
//				fmt.Println("Someee error::: ",err)
//			}
//
//			if err == io.EOF {
//				break
//				goto END
//			}
//
//			if filecreated == false {
//				filecreated = true
//				fo, err := os.Create(outputFile)
//				if err != nil {
//					panic(err)
//				}
//				// open output file
//				fo, err := os.Create("web/destImgs/"+data.FileName)
//
//				if err != nil {
//					panic(err)
//				}
//
//				// close fo on exit and check for its returned error
//				defer func() {
//					if err := fo.Close(); err != nil {
//						panic(err)
//					}
//				}()
//			}
//
//			_, writeErr := fo.Write(data.Content)
//			if writeErr != nil {
//				fmt.Println(data.Content)
//				fmt.Println("write err", writeErr)
//			}
//
//		}
//		END:
//			fmt.Println("End!!!!!")
//	}()
//	wg.Wait()
//
//
//	//response, err := client.GetPaths(context.Background(), &api.UploadStatus{
//	//	Message: "File uploaded successfully",
//	//});
//	//
//	//if err != nil {
//	//	log.Fatalf("There was an error GetPaths: ", err)
//	//}
//	//
//	//services.ImagePage = response
//	//
//	//return
//
//	return
//}
//
//
