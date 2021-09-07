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
 * 演示二管理
 * @author 半城风雨
 * @since 2021/08/07
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
        
             
			, {field: 'name', width: 100, title: '演示名称', align: 'center'}
            
        
            
            , {field: 'status', width: 100, title: '状态', align: 'center', templet: function (d) {
				return  '<input type="checkbox" name="status" value="' + d.id + '" lay-skin="switch" lay-text="正常|停用" lay-filter="Status" '+(d.status==1 ? 'checked' : '')+'>';
            }}
            
        
             
			, {field: 'sort', width: 100, title: '排序号', align: 'center'}
            
        
            , {field: 'createTime', width: 180, title: '添加时间', align: 'center'}
            , {field: 'updateTime', width: 180, title: '更新时间', align: 'center'}
            , {fixed: 'right', width: 150, title: '功能操作', align: 'center', toolbar: '#toolBar'}
        ];

        //【渲染TABLE】
        func.tableIns(cols, "tableList");

        //【设置弹框】
        
        func.setWin("演示二", 750, 450);
        

    
		
	
		
		//【设置状态】
        func.formSwitch('Status', null, function (data, res) {
            console.log("开关回调成功");
        });
		
	
		
	
    }
});
