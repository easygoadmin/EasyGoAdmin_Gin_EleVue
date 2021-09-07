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
 * 代码生成器
 * @author 半城风雨
 * @since 2021/7/26
 */
layui.use(['func'], function () {

    //【声明变量】
    var func = layui.func
        , $ = layui.$;

    if (A == 'index') {
        //【TABLE列数组】
        var cols = [
            {type: 'checkbox', fixed: 'left'}
            , {field: 'name', width: 150, title: '表名', align: 'center'}
            , {field: 'engine', width: 100, title: '引擎', align: 'center'}
            , {field: 'version', width: 100, title: '版本', align: 'center'}
            , {field: 'collation', width: 180, title: '编码', align: 'center'}
            , {field: 'rows', width: 100, title: '记录数', align: 'center'}
            , {field: 'data_length', width: 100, title: '大小', align: 'center'}
            , {field: 'auto_increment', width: 100, title: '自增索引', align: 'center'}
            , {field: 'comment', width: 150, title: '表备注', align: 'center'}
            , {
                field: '', width: 100, title: '状态', align: 'center', templet: function (d) {
                    return '未备份';
                }
            }
            , {field: 'createTime', width: 180, title: '创建时间', align: 'center', sort: true}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center', sort: true}
            , {fixed: 'right', width: 100, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList", function (layEvent, data) {
            if (layEvent === 'generate') {
                // 一键生成模块
                layer.confirm('您确定要生成表【' + data.name + '】对应的模块吗？生成后将会覆盖已有的模块文件！', {
                    icon: 3,
                    skin: 'layer-ext-moon',
                    btn: ['确认', '取消'] //按钮
                }, function (index) {
                    // layer.msg("一键生成模块");
                    // 调用内部方法
                    var url = cUrl + "/generate";
                    func.ajaxPost(url, {"name": data.name, "comment": data.comment}, function (data, flag) {
                        // 关闭弹窗
                        layer.close(index);
                    }, '模块文件生成中。。。');
                });
            }
        });

        // 批量生成
        $("#batchGenerate").click(function () {
            // 选择数据
            var data = func.getCheckData();
            if (data.length == 0) {
                layer.msg("请选择数据表", {icon: 5});
                return false;
            }
            layer.msg("批量生成功能研发中,请耐心等待...")
        });
    }
});
