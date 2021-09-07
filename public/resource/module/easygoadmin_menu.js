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
 * 菜单管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func', 'common', 'form', 'transfer'], function () {

    //声明变量
    var func = layui.func
        , common = layui.common
        , form = layui.form
        , transfer = layui.transfer
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {field: 'id', width: 80, title: 'ID', align: 'center', sort: true}
            , {field: 'name', width: 200, title: '菜单名称', align: 'left'}
            , {field: 'type', width: 80, title: '类型', align: 'center', templet(d) {
                    if (d.type == 0) {
                        // 菜单
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">菜单</span>';
                    } else if (d.type == 1) {
                        // 节点
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">节点</span>';
                    }
                }
            }
            , { field: 'icon', width: 80, title: '图标', align: 'center', templet: '<p><i class="layui-icon {{d.icon}}"></i></p>'}
            , {field: 'url', width: 150, title: 'URL地址', align: 'center'}
            , {field: 'permission', width: 180, title: '权限标识', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet(d) {
                    if (d.status == 1) {
                        // 在用
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">在用</span>';
                    } else {
                        // 停用
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">停用</span>';
                    }
                }
            }
            , {field: 'target', width: 100, title: '是否公共', align: 'center', templet(d) {
                    if (d.target == 1) {
                        // 内部打开
                        return '<span class="layui-btn layui-btn-normal layui-btn-xs">内部打开</span>';
                    } else if (d.target == 2) {
                        // 外部打开
                        return '<span class="layui-btn layui-btn-primary layui-btn-xs">外部打开</span>';
                    }
                }
            }
            , {field: 'sort', width: 90, title: '显示顺序', align: 'center'}
            , {fixed: 'right', width: 220, title: '功能操作', align: 'left', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.treetable(cols, "tableList");

        //【设置弹框】
        func.setWin("菜单");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });

    } else {

        // 初始化
        var type = $("#type").val()
        if (type == 0) {
            $(".func").removeClass("layui-hide");
        } else {
            $(".func").addClass("layui-hide");
        }

        // 菜单类型选择事件
        form.on('select(type)', function (data) {
            var val = data.value;
            if (val == 0) {
                $(".func").removeClass("layui-hide");
            } else {
                $(".func").addClass("layui-hide");
            }
        });

        /**
         * 提交表单
         */
        form.on('submit(submitForm2)', function (data) {
            if (data.field['type'] == 0) {
                // 获取穿梭组件的选中值
                var funcList = transfer.getData('func'); //获取右侧数据
                // 重组数据并赋值给字段
                var item = [];
                $.each(funcList, function (key, val) {
                    item.push(val['value']);
                });
                data.field['func'] = item.join(",");
            }
            // 提交表单
            common.submitForm(data.field, null, function (res, success) {
                console.log("保存成功回调");
            });
            return false;
        });
    }
});
