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
 * 配置管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['layer', 'form', 'table', 'util', 'admin', 'func'], function () {
    var $ = layui.jquery;
    var layer = layui.layer;
    var form = layui.form;
    var table = layui.table;
    var util = layui.util;
    var admin = layui.admin;
    var func = layui.func;
    var selObj;  // 左表选中数据

    form.on('submit(configDataEditSubmit)', function (data) {
        console.log("1111")
    });

    /* 渲染表格 */
    var insTb = table.render({
        elem: '#configTable',
        url: '/config/list',
        height: 'full-100',
        method: 'post',
        toolbar: ['<p>',
            '<button lay-event="add" class="layui-btn layui-btn-sm icon-btn"><i class="layui-icon">&#xe654;</i>添加</button>&nbsp;',
            '<button lay-event="edit" class="layui-btn layui-btn-sm layui-btn-warm icon-btn"><i class="layui-icon">&#xe642;</i>修改</button>&nbsp;',
            '<button lay-event="del" class="layui-btn layui-btn-sm layui-btn-danger icon-btn"><i class="layui-icon">&#xe640;</i>删除</button>',
            '</p>'].join(''),
        defaultToolbar: [],
        cols: [[
            {field: 'id', width: 80, title: 'ID', align: 'center'}
            , {field: 'name', title: '配置名称'}
        ]],
        done: function (res, curr, count) {
            $('#configTable+.layui-table-view .layui-table-body tbody>tr:first').trigger('click');
        }
    });

    /* 表格搜索 */
    form.on('submit(configTbSearch)', function (data) {
        insTb.reload({where: data.field});
        return false;
    });

    /* 表格头工具栏点击事件 */
    table.on('toolbar(configTable)', function (obj) {
        if (obj.event === 'add') { // 添加
            showEditModel();
        } else if (obj.event === 'edit') { // 修改
            showEditModel(selObj.data, selObj);
        } else if (obj.event === 'del') { // 删除
            doDel(selObj);
        }
    });

    /* 监听行单击事件 */
    table.on('row(configTable)', function (obj) {
        selObj = obj;
        obj.tr.addClass('layui-table-click').siblings().removeClass('layui-table-click');
        insTb2.reload({where: {configId: obj.data.id}, page: {curr: 1}, url: '/configdata/list'});
    });

    /* 显示表单弹窗 */
    function showEditModel(mData, obj) {
        admin.open({
            type: 1,
            title: (mData ? '修改' : '添加') + '配置',
            content: $('#configEditDialog').html(),
            success: function (layero, dIndex) {
                // 回显表单数据
                form.val('configEditForm', mData);
                // 表单提交事件
                form.on('submit(configEditSubmit)', function (data) {
                    var loadIndex = layer.load(2);
                    $.post(mData ? '/config/update' : '/config/add', data.field, function (res) {
                        layer.close(loadIndex);
                        if (0 === res.code) {
                            layer.close(dIndex);
                            layer.msg(res.msg, {icon: 1});
                            if (obj) {
                                obj.update(data.field);
                            } else {
                                insTb.reload();
                            }
                        } else {
                            layer.msg(res.msg, {icon: 2});
                        }
                    }, 'json');
                    return false;
                });
            }
        });
    }

    /* 删除 */
    function doDel(obj) {
        layer.confirm('确定要删除此配置吗？', {
            skin: 'layui-layer-admin',
            shade: .1
        }, function (i) {
            layer.close(i);
            var loadIndex = layer.load(2);
            $.post('/config/delete', {
                ids: obj.data.id,
            }, function (res) {
                layer.close(loadIndex);
                if (0 === res.code) {
                    layer.msg(res.msg, {icon: 1});
                    obj.del();
                    $('#configTable+.layui-table-view .layui-table-body tbody>tr:first').trigger('click');
                } else {
                    layer.msg(res.msg, {icon: 2});
                }
            }, 'json');
        });
    }

    /* 渲染表格2 */
    var insTb2 = table.render({
        elem: '#configDataTable',
        data: [],
        height: 'full-100',
        method: 'post',
        page: true,
        toolbar: ['<p>',
            '<button lay-event="add" class="layui-btn layui-btn-sm icon-btn"><i class="layui-icon">&#xe654;</i>添加</button>&nbsp;',
            '<button lay-event="del" class="layui-btn layui-btn-sm layui-btn-danger icon-btn"><i class="layui-icon">&#xe640;</i>删除</button>&nbsp;',
            '</p>'].join(''),
        cellMinWidth: 100,
        cols: [[
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'title', width: 150, title: '配置标题', align: 'center'}
            , {field: 'code', width: 150, title: '配置标签符', align: 'center'}
            , {field: 'value', width: 150, title: '配置值', align: 'center'}
            , {field: 'options', width: 100, title: '配置项', align: 'center'}
            , {field: 'type', width: 100, title: '配置类型', align: 'center', templet(d) {
                    var cls = "";
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.typeName + '</span>';
                }
            }
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: function (d) {
                return  '<input type="checkbox" name="status" value="' + d.id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" '+(d.status==1 ? 'checked' : '')+'>';
            }}
            , {field: 'sort', width: 100, title: '排序号', align: 'center'}
            , {field: 'note', width: 100, title: '配置说明', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {title: '操作', toolbar: '#configDataTbBar', align: 'center', width: 120, minWidth: 120, fixed: 'right'}
        ]]
    });

    /* 表格2搜索 */
    form.on('submit(configDataTbSearch)', function (data) {
        insTb2.reload({where: data.field, page: {curr: 1}});
        return false;
    });

    /* 表格2工具条点击事件 */
    table.on('tool(configDataTable)', function (obj) {
        if (obj.event === 'edit') { // 修改
            showEditModel2(obj.data);
        } else if (obj.event === 'del') { // 删除
            doDel2(obj);
        }
    });

    /* 表格2头工具栏点击事件 */
    table.on('toolbar(configDataTable)', function (obj) {
        if (obj.event === 'add') { // 添加
            showEditModel2();
        } else if (obj.event === 'del') { // 删除
            var checkRows = table.checkStatus('configDataTable');
            if (checkRows.data.length === 0) {
                layer.msg('请选择要删除的数据', {icon: 2});
                return;
            }
            var ids = checkRows.data.map(function (d) {
                return d.id;
            });
            doDel2({ids: ids});
        }
    });

    /* 显示表单弹窗2 */
    function showEditModel2(mData) {
        admin.open({
            type: 1,
            title: (mData ? '修改' : '添加') + '配置项',
            content: $('#configDataEditDialog').html(),
            area: ['750px', '530px'],
            success: function (layero, dIndex) {
                // 回显表单数据
                form.val('configDataEditForm', mData);
                // 表单提交事件
                form.on('submit(configDataEditSubmit)', function (data) {
                    data.field.configId = selObj.data.id;
                    var loadIndex = layer.load(2);
                    $.post(mData ? '/configdata/update' : '/configdata/add', data.field, function (res) {
                        layer.close(loadIndex);
                        if (0 === res.code) {
                            layer.close(dIndex);
                            layer.msg(res.msg, {icon: 1});
                            insTb2.reload({page: {curr: 1}});
                        } else {
                            layer.msg(res.msg, {icon: 2});
                        }
                    }, 'json');
                    return false;
                });
            }
        });
    }

    /* 删除2 */
    function doDel2(obj) {
        layer.confirm('确定要删除选中数据吗？', {
            skin: 'layui-layer-admin',
            shade: .1
        }, function (i) {
            layer.close(i);
            var loadIndex = layer.load(2);
            var ids = []
            if (obj.data) {
                ids = [obj.data.id]
            } else if (obj.ids) {
                ids = obj.ids;
            }
            $.post('/configdata/delete', {ids: ids.join(",")}, function (res) {
                layer.close(loadIndex);
                if (0 === res.code) {
                    layer.msg(res.msg, {icon: 1});
                    insTb2.reload({page: {curr: 1}});
                } else {
                    layer.msg(res.msg, {icon: 2});
                }
            }, 'json');
        });
    }

    //【设置状态】
    func.formSwitch('status', "/configdata/setStatus", function (data, res) {
        console.log("开关回调成功");
    });

});
