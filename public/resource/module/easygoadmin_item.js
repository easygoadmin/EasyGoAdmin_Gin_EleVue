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
 * 站点管理
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
            , {field: 'name', width: 200, title: '站点名称', align: 'center'}
            , {field: 'type', width: 100, title: '站点类型', align: 'center', templet(d) {
                var cls = "";
                if (d.type == 1) {
                    // 普通站点
                    cls = "layui-btn-normal";
                } else if (d.type == 2) {
                    // 其他
                    cls = "layui-btn-danger";
                }
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.typeName+'</span>';
            }}
            , {field: 'url', width: 200, title: '站点地址', align: 'center', templet(d) {
                    return "<a href='" + d.url + "' target='_blank'>" + d.url + "</a>";
                }
            }
            , {field: 'image', width: 100, title: '站点图片', align: 'center', templet: function (d) {
                if (d.image) {
                    return '<a href="' + d.image + '" target="_blank"><img src="' + d.image + '" height="26" /></a>';
                }
              }
            }
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: function (d) {
                  return d.status == 1 ? "在用" : "停用";
                }}
            , {field: 'note', width: 100, title: '站点备注', align: 'center'}
            , {field: 'sort', width: 100, title: '显示顺序', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("站点");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    }
});
