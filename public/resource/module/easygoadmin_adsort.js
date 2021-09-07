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
 * 广告位描述
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
            , {field: 'description', width: 200, title: '广告位描述', align: 'center'}
            , {field: 'itemName', width: 150, title: '所属站点', align: 'center'}
            , {field: 'cateName', width: 200, title: '所属栏目', align: 'center'}
            , {field: 'locId', width: 120, title: '广告页面位置', align: 'center'}
            , {field: 'platform', width: 100, title: '所属平台', align: 'center', templet(d) {
                    var cls = "";
                    if (d.platform == 1) {
                        // PC网站
                        cls = "layui-btn-normal";
                    } else if (d.platform == 2) {
                        // WAP手机站
                        cls = "layui-btn-danger";
                    } else if (d.platform == 3) {
                        // 微信小程序
                        cls = "layui-btn-warm";
                    } else if (d.platform == 4) {
                        // APP移动端
                        cls = "layui-btn-primary";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.platformName + '</span>';
                }
            }
            , {field: 'sort', width: 100, title: '排序号', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("广告位描述",750, 400);

    }
});
