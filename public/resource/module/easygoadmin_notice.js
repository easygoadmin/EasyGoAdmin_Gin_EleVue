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
 * 通知公告
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        , form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
              {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'title', width: 350, title: '通知标题', align: 'center'}
            , {field: 'source', width: 100, title: '通知来源', align: 'center', templet(d) {
                var cls = "";
                if (d.source == 1) {
                    cls = "layui-btn-normal";
                } else if (d.source == 2) {
                    cls = "layui-btn-primary"
                }
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.sourceName+'</span>';
            }}
            , {field: 'isTop', width: 100, title: '是否置顶', align: 'center', templet(d) {
                if (d.isTop == 1) {
                    // 已置顶
                    return '<span class="layui-btn layui-btn-primary layui-btn-xs">已置顶</span>';
                } else if (d.isTop == 2) {
                    // 未置顶
                    return '<span class="layui-btn layui-btn-danger layui-btn-xs">未置顶</span>';
                }
            }}
            , {field: 'status', width: 100, title: '发布状态', align: 'center', templet(d) {
                if (d.status == 1) {
                    // 草稿箱
                    return '<span class="layui-btn layui-btn-normal layui-btn-xs">已发布</span>';
                } else if (d.status == 2) {
                    // 立即发布
                    return '<span class="layui-btn layui-btn-danger layui-btn-xs">待发布</span>';
                }
            }}
            , {field: 'browse', width: 100, title: '阅读量', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("通知公告");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
