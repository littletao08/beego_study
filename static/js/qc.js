$(function () {
    //页面加载完成以后获取用户的QQ信息
    $(document).on("load", function () {
        if (QC.Login.check()) {
            QC.Login.getMe(function (openId, accessToken) {
                // 拿到用户的OPENID 和token
                // 发送请求到后台,如果这个用户不存在,跳转到用户设置密码页面.
                // 如果这个用户存在,则保证这个请求是正常从QQ登录过来的,就可以
                var qcUrl ="https://graph.qq.com/user/get_user_info"
                var qcParams = {access_token:accessToken,openid:openId,oauth_consumer_key:"101284166"};
                var nickname = "";
                var sex = "";
                $.get(qcUrl,qcParams,function(resoponse){
                    nickname = QC.String.escHTML(resoponse.nickname)
                    sex = response.gender;
                },function(){
                    //ignore 没有获取到用户的QQ登录信息.
                    window.location="/index";
                })
                var data = {name:nickname,qcOpenId:openId,sex:sex};
                $.post("/users/qclogin", data, function (response) {
                    window.location="/index";
                },function(){
                    window.location="/index";
                })
            })
        }
    })
})