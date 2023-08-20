<template>
  <div>
    <div ref="texteditor"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, createApp } from "vue";
import MyDialogContent from "./MyDialogContent.vue";
let content = ref("");
let texteditor = ref(null);

// 监听内容变化
onMounted(() => {
  initTinyMCE();
});
// 卸载前销毁编辑器
onBeforeUnmount(() => {
  if (window.tinymce) {
    window.tinymce.EditorManager.execCommand(
      "mceRemoveEditor",
      true,
      texteditor.value.id
    );
  }
});

// 初始化 TinyMCE
const initTinyMCE = () => {
  if (process.client) {
    if (window.tinymce) {
      initializeEditor();
      return;
    }

    // 使用 ID 检查是否已经加载了 TinyMCE 脚本
    const existingScript = document.getElementById("tinymce-script");

    if (!existingScript) {
      const script = document.createElement("script");
      script.id = "tinymce-script"; // 给 script 标签添加 ID
      script.src = "/tinymce/js/tinymce/tinymce.min.js";
      script.onload = initializeEditor;
      document.body.appendChild(script);
    } else {
      // 如果 script 已存在，但 window.tinymce 仍未定义，那么可能是 script 还在加载中。
      // 使用事件监听器来检测何时 script 加载完成。
      existingScript.addEventListener("load", initializeEditor);
    }
  }
};
// 初始化编辑器
const initializeEditor = () => {
  window.tinymce.init({
    target: texteditor.value,
    height: 500,
    language: "zh-Hans",
    statusbar: false, // 状态栏
    plugins: [
      "advlist",
      "anchor",
      "autolink",
      "charmap",
      "code",
      "fullscreen",
      "help",
      "image",
      "insertdatetime",
      "link",
      "lists",
      "media",
      "preview",
      "searchreplace",
      "table",
      "visualblocks",
      "accordion",
      "quickbars",
      "preview",
    ],
    skin: "oxide", // 使用 oxide 主题
    toolbar:
      "undo redo | bold italic underline | emoticons fontselect fontsizeselect forecolor backcolor | align accordion | numlist bullist outdent indent | link unlink uploadimages image media table | blockquote hr removeformat | preview fullscreen", // 添加 'image' 按钮到工具栏
    menubar: false,
    branding: false,
    quickbars_image_toolbar: "rotateimage",
    // 上传图片
    images_upload_handler: customImageUploadHandler,
    setup: registerEventHandlers,
  });
};

/**
 * 注册事件处理器
 *
 * @param {tinymce.Editor} editor
 * @see https://www.tiny.cloud/docs/advanced/events/
 */
function registerEventHandlers(editor) {
  editor.ui.registry.addButton("rotateimage", {
    text: "",
    icon: "rotate-left", // 使用 TinyMCE 的内置图标
    onAction: function () {
      const selectedNode = editor.selection.getNode();
      if (selectedNode && selectedNode.tagName === "IMG") {
        const currentRotation =
          parseInt(selectedNode.style.transform.replace(/[^0-9]/g, "")) || 0;
        const newRotation = (currentRotation + 90) % 360;
        editor.dom.setStyle(
          selectedNode,
          "transform",
          `rotate(${newRotation}deg)`
        );
      }
    },
  });
  editor.ui.registry.addButton("uploadimages", {
    text: "Upload Images",
    onAction: function () {
      openImageUploadDialog(editor);
    },
  });
}
async function getDialogContentHtml() {
  // 创建一个临时的 DOM 元素
  const container = document.createElement("div");

  // 使用 createApp 方法挂载 Vue 组件到这个临时的 DOM 元素
  const app = createApp(MyDialogContent);
  app.mount(container);

  // 返回 HTML 内容
  return container.innerHTML;
}

/**
 * 打开图片上传对话框
 *
 * @param {tinymce.Editor} editor
 */
async function openImageUploadDialog(editor) {
  const contentHtml = await getDialogContentHtml();
  editor.windowManager.open({
    title: "Upload Multiple Images",
    body: {
      type: "panel",
      items: [
        {
          type: "htmlpanel",
          html: contentHtml,
        },
      ],
    },
    buttons: [
      {
        type: "cancel",
        text: "Cancel",
      },
      {
        type: "submit",
        text: "Insert",
        primary: true,
      },
    ],
    onSubmit: function (api) {
      const data = api.getData();
      // 这里你可以处理图片上传逻辑
      // 并在完成后使用 editor.insertContent 插入图片到编辑器
      api.close();
    },
  });
}

/**
 * 自定义图片上传处理器
 *
 * @param {BlobInfo} blobInfo
 * @param {function} progress
 * @returns {Promise}
 * @see https://www.tiny.cloud/docs/configure/file-image-upload/#images_upload_handler
 */
const customImageUploadHandler = (blobInfo, progress) =>
  new Promise((resolve, reject) => {
    // 创建 XMLHttpRequest 对象
    const xhr = new XMLHttpRequest();

    // 初始化 POST 请求，并指定上传处理器 URL
    xhr.open("POST", "postAcceptor.php");

    // 监听上传进度并报告
    xhr.upload.onprogress = (e) => progress((e.loaded / e.total) * 100);

    // 当请求完成时处理响应
    xhr.onload = () => {
      resolve("https://images.hxsj.in/test/1.jpg");
      // if (xhr.status >= 200 && xhr.status < 300) {
      //   const json = JSON.parse(xhr.responseText);
      //   if (json && typeof json.location === "string") {
      //     resolve(json.location);
      //   } else {
      //     reject("无效的 JSON 响应: " + xhr.responseText);
      //   }
      // } else {
      //   reject("HTTP 错误: " + xhr.status);
      // }
    };

    // 请求失败时的处理
    xhr.onerror = () => reject("XHR 请求失败，错误码: " + xhr.status);

    // 构建 FormData 对象并添加要上传的文件
    const formData = new FormData();
    formData.append("file", blobInfo.blob(), blobInfo.filename());

    // 发送请求
    xhr.send(formData);
  });
</script>
