var wrapper = $('.wrapper');
var _time = null;

$(function () {
    clearTimeout(_time);

    var drop = $("input");
    var alertInfo = $(".alert-info");

    drop.on('dragenter', function (e) {
        wrapper.addClass("active");

    }).on('dragleave dragend mouseout drop', function (e) {
        wrapper.removeClass("active");
    });

    $('#uploadImage').change(handleFileSelect);
    $(".btn-clear").click(clearInput);

    if (alertInfo.is(":visible")) {
        _time = setTimeout(function () {
            alertInfo.fadeOut(200);
        }, 1500)
    }
});

function handleFileSelect(evt) {
    var file = evt.target.files[0]; // FileList object

    // Only process image files.
    if (file && !file.type.match('image.*')) {
        return;
    }

    var reader = new FileReader();

    // Closure to capture the file information.
    reader.onload = (function(theFile) {
        return function(e) {
            previewImage(wrapper, e.target.result)
        };
    })(file);

    // Read in the image file as a data URL.
    reader.readAsDataURL(file);
}

function previewImage(container, imgSrc) {
    var imgPreview = $("#image-preview");

    imgPreview.attr("src", imgSrc);
    imgPreview.removeClass("hidden");
    $('.btn-upload, .btn-clear').removeClass("hidden");

    $(".wrapper").hide()
}

function clearInput() {
    var imgPreview = $("#image-preview");

    imgPreview.attr("src", "");
    imgPreview.addClass("hidden");
    $('.btn-upload, .btn-clear').addClass("hidden");
    $(".wrapper").show()
}