/**
 * Created by zhanglida on 15/11/24.
 */
$(function () {

    var $window = $(window)
    //边栏滚动控制

    //给赞助添加事件
    $("#sponsor").click(function(){
        $.post("/sponsors/new",function(data){
            $(data).submit();
        })
    });

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

