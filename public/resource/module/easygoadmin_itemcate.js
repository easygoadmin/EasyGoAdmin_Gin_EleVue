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
 * 栏目管理
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
              {field: 'id', width: 80, title: 'ID', align: 'center', sort: true}
            , {field: 'name', width: 200, title: '栏目名称', align: 'left'}
            , {field: 'itemName', width: 200, title: '所属站点', align: 'center'}
            , {field: 'pinyin', width: 150, title: '拼音(全)', align: 'center'}
            , {field: 'code', width: 100, title: '拼音(简)', align: 'center'}
            , {field: 'is_cover', width: 100, title: '有无封面', align: 'center', templet(d) {
                    if (d.is_cover == 1) {
                        // 有封面
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">有封面</span>';
                    } else {
                        // 无封面
                        return '<span class="layui-btn layui-btn-danger layui-btn-xs">有封面</span>';
                    }
                }}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet(d) {
                if (d.status == 1) {
                    // 在用
                    return '<span class="layui-btn layui-btn-normal layui-btn-xs">在用</span>';
                } else {
                    // 停用
                    return '<span class="layui-btn layui-btn-danger layui-btn-xs">停用</span>';
                }
            }}
            , {field: 'sort', width: 100, title: '排序号', align: 'center'}
            , {field: 'note', width: 200, title: '备注', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {width: 220, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.treetable(cols, "tableList");

        //【设置弹框】
        func.setWin("栏目");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
