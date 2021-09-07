/**
 * 上传图片裁剪组件
 * @author 牧羊人
 * @date 2018/9/13
 */
layui.extend({
    cropper: '/lib/extend/cropper/cropper',
}).define(['jquery', 'layer', 'cropper'], function (exports) {
    var $ = layui.$
        , layer = layui.layer;

    var obj = {
        render: function (e) {

            var self = this,
                elem = e.elem,
                name = e.name,
                saveW = e.saveW,
                saveH = e.saveH,
                mark = e.mark,
                area = e.area,
                url = e.url,
                done = e.done;

            var html = "<link rel=\"stylesheet\" href=\"/static/lib/extend/cropper/cropper.css\">\n" +
                "<div class=\"layui-fluid showImgEdit_" + name + "\" style=\"display: none\">\n" +
                "    <div class=\"layui-form-item\" style=\"margin-top:0px;\">\n" +
                "        <div class=\"layui-input-inline layui-btn-container\" style=\"width: auto;\">\n" +
                "            <label for=\"cropper_avatarImgUpload_" + name + "\" class=\"layui-btn layui-btn-primary\">\n" +
                "                <i class=\"layui-icon\">&#xe67c;</i>选择图片\n" +
                "            </label>\n" +
                "            <input class=\"layui-upload-file\" id=\"cropper_avatarImgUpload_" + name + "\" type=\"file\" value=\"选择图片\" name=\"file\">\n" +
                "        </div>\n" +
                "        <div class=\"layui-form-mid layui-word-aux\">上传图片大小限定在500kb以内</div>\n" +
                "    </div>\n" +
                "    <div class=\"layui-row layui-col-space15\">\n" +
                "        <div class=\"layui-col-xs9\">\n" +
                "            <div class=\"readyimg\" style=\"height:300px;background-color: rgb(247, 247, 247);\">\n" +
                "                <img src=\"\" >\n" +
                "            </div>\n" +
                "        </div>\n" +
                "        <div class=\"layui-col-xs3\">\n" +
                "            <div class=\"img-preview\" style=\"width:150px;height:150px;overflow:hidden\">\n" +
                "            </div>\n" +
                "        </div>\n" +
                "    </div>\n" +
                "    <div class=\"layui-row layui-col-space15\">\n" +
                "        <div class=\"layui-col-xs9\">\n" +
                "            <div class=\"layui-row\">\n" +
                "                <div class=\"layui-col-xs6\">\n" +
                "                    <button type=\"button\" class=\"layui-btn layui-icon layui-icon-left\" cropper-event=\"rotate\" data-option=\"-15\" title=\"Rotate -90 degrees\"> 向左旋转</button>\n" +
                "                    <button type=\"button\" class=\"layui-btn layui-icon layui-icon-right\" cropper-event=\"rotate\" data-option=\"15\" title=\"Rotate 90 degrees\"> 向右旋转</button>\n" +
                "                </div>\n" +
                // "                <div class=\"layui-col-xs5\" style=\"text-align: right;\">\n" +
                // "                    <button type=\"button\" class=\"layui-btn layui-icon layui-icon-set-fill\" title=\"移动\"></button>\n" +
                // "                    <button type=\"button\" class=\"layui-btn layui-icon layui-icon-share\" title=\"放大图片\"></button>\n" +
                // "                    <button type=\"button\" class=\"layui-btn layui-icon layui-icon-share\" title=\"缩小图片\"></button>\n" +
                // "                    <button type=\"button\" class=\"layui-btn layui-icon layui-icon-refresh\" cropper-event=\"reset\" title=\"重置图片\"></button>\n" +
                // "                </div>\n" +
                "            </div>\n" +
                "        </div>\n" +
                "        <div class=\"layui-col-xs3\">\n" +
                "            <button class=\"layui-btn layui-btn-fluid\" cropper-event=\"confirmSave\" type=\"button\"> 保存修改</button>\n" +
                "        </div>\n" +
                "    </div>\n" +
                "\n" +
                "</div>";
            $('body').append(html);

            var content = $('.showImgEdit_' + name)
                , image = $(".showImgEdit_" + name + " .readyimg img")
                , preview = '.showImgEdit_' + name + ' .img-preview'
                , file = $(".showImgEdit_" + name + " input[name='file']")
                , options = {aspectRatio: mark, preview: preview, viewMode: 1};

            $(elem).on('click', function () {
                layer.open({
                    title: '选择图片裁剪'
                    , type: 1
                    , content: content
                    , area: area
                    , success: function () {
                        image.cropper(options);
                    }
                    , cancel: function (index) {
                        layer.close(index);
                        image.cropper('destroy');
                    }
                });
                return false;
            });
            $(".layui-btn").on('click', function () {
                var event = $(this).attr("cropper-event");
                //监听确认保存图像
                if (event === 'confirmSave') {
                    image.cropper("getCroppedCanvas", {
                        width: saveW,
                        height: saveH
                    }).toBlob(function (blob) {
                        var formData = new FormData();
                        formData.append('file', blob, 'head.jpg');

                        var index = layer.msg('图片保存中，请稍候', {icon: 16, time: false, shade: 0.2});

                        $.ajax({
                            method: "post",
                            url: url, //用于文件上传的服务器端请求地址
                            data: formData,
                            processData: false,
                            contentType: false,
                            success: function (result) {
                                if (result.code == 0) {
                                    //关闭提示层
                                    layer.close(index);
                                    layer.msg(result.msg, {icon: 1});
                                    layer.closeAll('page');
                                    return done(result.data);
                                } else {
                                    layer.close(index);
                                    layer.alert(result.msg, {icon: 2});
                                    return false;
                                }

                            }
                        });
                    });
                    //监听旋转
                } else if (event === 'rotate') {
                    var option = $(this).attr('data-option');
                    image.cropper('rotate', option);
                    //重设图片
                } else if (event === 'reset') {
                    image.cropper('reset');
                }
                //文件选择
                file.change(function () {
                    var r = new FileReader();
                    var f = this.files[0];
                    r.readAsDataURL(f);
                    r.onload = function (e) {
                        image.cropper('destroy').attr('src', this.result).cropper(options);
                    };
                });
            });
        }

    };
    exports('croppers', obj);
});
