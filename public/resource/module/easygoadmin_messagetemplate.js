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
 * 消息模板
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
              {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'code', width: 200, title: '模板CODE', align: 'center'}
            , {field: 'title', width: 200, title: '模板标题', align: 'center'}
            , {field: 'type', width: 100, title: '模板类型', align: 'center', templet(d) {
                var cls = "";
                if (d.type == 1) {
                    // 系统模板
                    cls = "layui-btn-normal";
                } else if (d.type == 2) {
                    // 短信模板
                    cls = "layui-btn-danger";
                } else if (d.type == 3) {
                    // 邮件模板
                    cls = "layui-btn-warm";
                } else if (d.type == 4) {
                    // 微信模板
                    cls = "layui-btn-primary";
                } else if (d.type == 5) {
                    // 推送模板
                    cls = "layui-btn-disabled";
                }

				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.typeName+'</span>';
            }}
            , {field: 'content', width: 250, title: '模板内容', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: '#statusTpl'}
            , {field: 'createUserName', width: 100, title: '添加人', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateUserName', width: 100, title: '更新人', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("消息模板", 650, 480);

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
