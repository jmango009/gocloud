var isIE = /msie/i.test(navigator.userAgent) && !window.opera;

function imageChange(target,id) {
    clearPct()
    var filetypes =[".jpg",".png",".bmp",".jpeg"];
    var filemaxsize = 1024*10;//10M
    fileChange(target,filetypes,filemaxsize);
}

function videoChange(target,id) {
    clearPct()
    $('#file-list').html("");
    var filetypes =[".mp4"];
    var filemaxsize = 1024*50;//50M
    fileChange(target,filetypes,filemaxsize);
}

function fileChange(target,filetypes,filemaxsize) {
    var fileSize = 0;
    var filepath = target.value;
    if(filepath){
        var isnext = false;
        var fileend = filepath.substring(filepath.indexOf("."));
        if(filetypes && filetypes.length>0){
            for(var i =0; i<filetypes.length;i++){
                if(filetypes[i]==fileend){
                    isnext = true;
                    break;
                }
            }
        }
        if(!isnext){
            alert("不接受此文件类型！");
            target.value ="";
            return false;
        }
    }else{
        return false;
    }
    if (isIE && !target.files) {
        var filePath = target.value;
        var fileSystem = new ActiveXObject("Scripting.FileSystemObject");
        if(!fileSystem.FileExists(filePath)){
            alert("附件不存在，请重新输入！");
            return false;
        }
        var file = fileSystem.GetFile (filePath);
        fileSize = file.Size;
    } else {
        fileSize = target.files[0].size;
    }

    var size = fileSize / 1024;
    if(size>filemaxsize){
        alert("附件大小不能大于"+filemaxsize/1024+"M！");
        target.value ="";
        return false;
    }
    if(size<=0){
        alert("附件大小不能为0M！");
        target.value ="";
        return false;
    }
}

function clearPct() {
    $(".am-progress").hide()
    $(".am-progress-bar").css('width', 0);
    $(".am-progress-bar").html("");
}

function uploadVideo() {
    var video = $("#video").get(0).files[0];
    uploadFile(video, "video", "/video_list.html")
}

function uploadImage() {
    var pic = $("#pic").get(0).files[0];
    uploadFile(pic, "image", "/image_list.html")
}

function uploadFile(file, type, redirectUrl) {
    var formData = new FormData();
    formData.append("file", file);
    formData.append("description", $("#description").html());

    $.ajax({
        type: "POST",
        url: "/" + type,
        data: formData,
        processData: false,
        contentType: false,
        xhr: function () {
            var xhr = $.ajaxSettings.xhr();
            if (onprogress && xhr.upload) {
                xhr.upload.addEventListener("progress", onprogress, false);
                return xhr;
            }
        },
        success: function (data, textStatus) {
            $("#son").html("successfully");
            window.location.href = redirectUrl
        },
        error: function (xhr, textStatus, errorThrown) {
            $("#son").html(textStatus);
        }
    });
}

function onprogress(evt) {
    var loaded = evt.loaded;
    var tot = evt.total;
    var per = Math.floor(100 * loaded / tot);
    $(".am-progress").show()
    $(".am-progress-bar").css('width', per + "%");
    $(".am-progress-bar").html(per + "%");
}