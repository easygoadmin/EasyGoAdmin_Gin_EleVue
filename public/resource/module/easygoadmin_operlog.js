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
 * 操作日志
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
            , {field: 'title', width: 100, title: '模块标题', align: 'center'}
            , {field: 'businessType', width: 100, title: '业务类型', align: 'center', templet(d) {
                var cls = "";
                if (d.businessType == 1) {
                    // 新增
                    cls = "layui-btn-normal";
                } else if (d.businessType == 2) {
                    // 修改
                    cls = "layui-btn-danger";
                } else if (d.businessType == 3) {
                    // 删除
                    cls = "layui-btn-warm";
                } 
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.businessTypeName+'</span>';
            }}
            , {field: 'method', width: 100, title: '方法名称', align: 'center'}
            , {field: 'requestMethod', width: 100, title: '请求方式', align: 'center'}
            , {field: 'operatorType', width: 100, title: '操作类别', align: 'center', templet(d) {
                var cls = "";
                if (d.operatorType == 1) {
                    // 后台用户
                    cls = "layui-btn-normal";
                } else if (d.operatorType == 2) {
                    // 手机端用户
                    cls = "layui-btn-danger";
                } 
				return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.operatorTypeName+'</span>';
            }}
            , {field: 'operName', width: 100, title: '操作人员', align: 'center'}
            , {field: 'operUrl', width: 100, title: '请求URL', align: 'center'}
            , {field: 'operIp', width: 130, title: '主机地址', align: 'center'}
            , {field: 'operLocation', width: 100, title: '操作地点', align: 'center'}
            , {field: 'operParam', width: 100, title: '请求参数', align: 'center'}
            , {field: 'jsonResult', width: 100, title: '返回参数', align: 'center'}
            , {field: 'status', width: 100, title: '状态', align: 'center', templet(d) {
                    var cls = "";
                    if (d.status == 1) {
                        // 正常
                        cls = "layui-btn-normal";
                    } else if (d.status == 2) {
                        // 异常
                        cls = "layui-btn-danger";
                    }
                    return '<span class="layui-btn ' + cls + ' layui-btn-xs">'+d.statusName+'</span>';
                }}
            , {field: 'errorMsg', width: 100, title: '错误消息', align: 'center'}
            , {field: 'createTime', width: 180, title: '创建时间', align: 'center'}
            , {fixed: 'right', width: 100, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

    }
});
