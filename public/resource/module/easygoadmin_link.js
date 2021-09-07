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
 * 友链管理
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //声明变量
    var func = layui.func
        ,form = layui.form
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'id', width: 80, title: 'ID', align: 'center', sort: true, fixed: 'left'}
            , {field: 'name', width: 250, title: '友链名称', align: 'center'}
            , {field: 'image', width: 100, title: '友链图片', align: 'center', templet: function (d) {
                    if (d.image != "") {
                        return '<a href="' + d.image + '" target="_blank"><img src="' + d.image + '" height="26" /></a>';
                    }
                }
            }
            , {field: 'type', width: 100, title: '类型', align: 'center', templet(d) {
                    var cls = "";
                    if (d.type == 1) {
                        // 友情链接
                        cls = "layui-btn-normal";
                    } else if (d.type == 2) {
                        // 合作伙伴
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.typeName + '</span>';
                }
            }
            , {field: 'url', width: 200, title: '友链地址', align: 'center', templet(d) {
                    return "<a href='" + d.url + "' target='_blank'>" + d.url + "</a>";
                }
            }
            , {field: 'platform', width: 100, title: '平台', align: 'center', templet(d) {
                    var cls = "";
                    if (d.platform == 1) {
                        // PC站
                        cls = "layui-btn-normal";
                    } else if (d.platform == 2) {
                        // WAP站
                        cls = "layui-btn-danger";
                    } else if (d.platform == 3) {
                        // 微信小程序
                        cls = "layui-btn-warm";
                    } else if (d.platform == 4) {
                        // APP应用
                        cls = "layui-btn-primary";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.platformName + '</span>';
                }
            }
            , {field: 'form', width: 100, title: '友链形式', align: 'center', templet(d) {
                    var cls = "";
                    if (d.form == 1) {
                        // 文字链接
                        cls = "layui-btn-normal";
                    } else if (d.form == 2) {
                        // 图片链接
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">' + d.formName + '</span>';
                }
            }
            , {field: 'image', width: 100, title: '友链图片', align: 'center', templet: function (d) {
                    var imageStr = "";
                    if (d.imageUrl) {
                        imageStr = '<a href="' + d.imageUrl + '" target="_blank"><img src="' + d.imageUrl + '" height="26" /></a>';
                    }
                    return imageStr;
                }
            }
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '<input type="checkbox" name="status" value="' + d.id + '" lay-skin="switch" lay-text="正常|禁用" lay-filter="status" ' + (d.status == 1 ? 'checked' : '') + '>';
                }
            }
            , {field: 'sort', width: 100, title: '显示顺序', align: 'center'}
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        func.setWin("友链");

        //【设置状态】
        func.formSwitch('status', null, function (data, res) {
            console.log("开关回调成功");
        });
    } else {
        //【监听友链类型】
        var link_form = $('#form').val();
        if (link_form == 1) {
            //文字
            $(".image").addClass("layui-hide");
        } else if (link_form == 2) {
            //图片
            $(".image").removeClass("layui-hide");
        }
        form.on('select(form)', function (data) {
            if (data.value == 1) {
                $(".image").addClass("layui-hide");
            } else if (data.value == 2) {
                $(".image").removeClass("layui-hide");
            }
        });
    }
});
