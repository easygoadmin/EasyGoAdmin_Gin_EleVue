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
 * 个人中心
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['form', 'element', 'admin', 'func'], function () {
    var $ = layui.jquery;
    var form = layui.form;
    var element = layui.element;
    var admin = layui.admin;
    var func = layui.func;

    /* 选择头像 */
    $('#userInfoHead').click(function () {
        layer.msg("头像裁剪完善中");
        return false;
        // admin.cropImg({
        //     imgSrc: $('#userInfoHead>img').attr('src'),
        //     onCrop: function (res) {
        //         $('#userInfoHead>img').attr('src', res);
        //         parent.layui.jquery('.layui-layout-admin>.layui-header .layui-nav img.layui-nav-img').attr('src', res);
        //     }
        // });
    });

    /* 监听表单提交 */
    form.on('submit(userInfoSubmit)', function (data) {
        func.ajaxPost("/userInfo", data.field);
        return false;
    });

});
