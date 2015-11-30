/**
 * Created by zhanglida on 15/11/24.
 */
$(function () {

    var $window = $(window)

    //给赞助添加事件
    $("#sponsor").click(function(){
        $.post("/sponsors/new",function(data){
            $(data).submit();
        })
    });
    //左边栏滚动控制
    setTimeout(function () {
        $('.bs-docs-sidenav').affix({
            offset: {
                top: function () { return $window.width() <= 980 ? 290 : 220 }
                ,
               bottom: 270
            }
        })
    }, 100)

})


function initMarkdownEditor(){

    $("#editor").markdownEditor(
        {
            preview: true,
            onPreview: function (content, callback) {
                callback(marked(content));
            },
            onSave: function (content) {
                var content = marked(content);
                if (!$.trim(content)) {
                    $(".md-editor").css("border", "1px solid red");
                    return;
                }
                $.post('/articles', {title: "go study", content: content},
                    function (data) {
                        if (data) {
                            alert("保存成功!");
                            window.location.href = "/"
                        } else {
                            alert("保存失败!")
                        }
                    }
                )
            }
        }
    );

    $("textarea").keyup(function () {
        var content = $('#editor').markdownEditor('content');
        if (null != content && content.length > 0) {
            $(".md-editor").css("border", "");
        }else{
            $(".md-editor").css("border", "1px solid red");
        }
    });
}

