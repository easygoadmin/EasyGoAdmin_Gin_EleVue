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
 * 消息管理
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
            , {field: 'title', width: 200, title: '消息标题', align: 'center'}
            , {field: 'content', width: 300, title: '消息内容', align: 'center'}
            , {field: 'receiver', width: 100, title: '接收人', align: 'center'}
            , {field: 'type', width: 100, title: '发送方式', align: 'center', templet(d) {
                var cls = "";
                if (d.type == 1) {
                    // 系统
                    cls = "layui-btn-normal";
                } else if (d.type == 2) {
                    // 短信
                    cls = "layui-btn-danger";
                } else if (d.type == 3) {
                    // 邮件
                    cls = "layui-btn-warm";
                } else if (d.type == 4) {
                    // 微信
                    cls = "layui-btn-primary";
                } else if (d.type == 5) {
                    // 推送
                    cls = "layui-btn-disabled";
                }

				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.typeName+'</span>';
            }}
            , {field: 'sendTime', width: 180, title: '发送时间', align: 'center'}
            , {field: 'sendStatus', width: 100, title: '发送状态', align: 'center', templet(d) {
                var cls = "";
                if (d.sendStatus == 1) {
                    // 已发送
                    cls = "layui-btn-normal";
                } else if (d.sendStatus == 2) {
                    // 未发送
                    cls = "layui-btn-danger";
                } 
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.sendStatusName+'</span>';
            }}
            , {field: 'sendNum', width: 100, title: '发送次数', align: 'center'}
            , {field: 'note', width: 200, title: '发送备注', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 100, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("系统消息");

    }
});
