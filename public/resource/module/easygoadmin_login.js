// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ EasyGoAdmin ]
// +----------------------------------------------------------------------
// | 版权所有 2021 EasyGoAdmin深圳研发中心
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: 半城风雨 <easygoadmin@163.com>
// +----------------------------------------------------------------------

/**
 * 登录
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['layer', 'form', 'index'], function () {
    var $ = layui.jquery;
    var layer = layui.layer;
    var form = layui.form;
    var index = layui.index;
    $('.login-wrapper').removeClass('layui-hide');

    // 登录事件
    form.on('submit(loginSubmit)', function (data) {
        // 设置按钮文字“登录中...”及禁止点击状态
        $(data.elem).attr('disabled', true).text('登录中。。。');

        // 网络请求
        var loadIndex = layer.load(2);
        $.ajax({
            type: "POST",
            url: '/login',
            data: JSON.stringify(data.field),
            contentType: "application/json",
            dataType: "json",
            beforeSend: function () {
                // TODO...
            },
            success: function (res) {
                layer.close(loadIndex);
                if (res.code == 0) {
                    // 清除Tab记忆
                    index.clearTabCache();

                    // 设置登录成功状态
                    $(data.elem).attr('disabled', true).text('登录成功');

                    // 提示语
                    layer.msg('登录成功', {
                        icon: 1,
                        time: 1500
                    });

                    // 延迟3秒
                    setTimeout(function () {
                        // 跳转后台首页
                        window.location.href = "/index";
                    }, 2000);

                    return false;
                } else {
                    // 错误信息
                    layer.msg(res.msg, {icon: 2, anim: 6});
                    // 刷新验证码
                    $('img.login-captcha').click(function () {
                        this.src = '/captcha?t=' + (new Date).getTime();
                    }).trigger('click');

                    // 延迟3秒恢复可登录状态
                    setTimeout(function () {
                        // 设置按钮状态为登录”
                        var login_text = $(data.elem).text().replace('中。。。', '');
                        // 设置按钮为可点击状态
                        $(data.elem).text(login_text).removeAttr('disabled');
                    }, 1000);
                }
            },
            error: function () {
                layer.msg("AJAX请求异常");
            }
        });
        return false;
    });

    // 获取图片验证码
    $('img.login-captcha').click(function () {
        var url = "/captcha?t=" + (new Date).getTime();
        $.ajax({
            type: "get",
            url: url,
            success: function (res) {
                if (res.code == 0) {
                    this.src = res.data;
                    $("#imgcode").attr("src", res.data);
                    $("#idkey").val(res.idkey);
                }
            }
        });
    }).trigger('click');

});