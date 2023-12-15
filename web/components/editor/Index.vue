<template>
  <div class="shadow-center rounded-xl">
    <Skeleton :lines="18" v-if="show" />
    <div ref="texteditor"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from "vue";
import { upload } from "~/api/upload";

useHead({
  script: {
    src: "/tinymce/js/tinymce/tinymce.min.js",
  },
});
const {getPath} = usePath();
// const { sleep } = useSleep();
let { placeholder, moduleValue } = defineProps({
  moduleValue: {
    type: String,
    default: "",
  },
  placeholder: {
    type: String,
    default: "请输入内容",
  },
});
let emit = defineEmits(["update:moduleValue"]);
let show = ref(true);
let texteditor = ref(null);
let editorInstance;
// 监听内容变化
onMounted(async () => {
  initTinyMCE();
  // await sleep(500);
  // initializeEditor();
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
    // 如果 TinyMCE 已经加载，那么直接初始化编辑器 如果没有重试3次
    if (window.tinymce) {
      initializeEditor();
      return;
    } else {
      let count = 0;
      let timer = setInterval(() => {
        if (window.tinymce) {
          initializeEditor();
          clearInterval(timer);
        } else {
          count++;
          if (count > 3) {
            clearInterval(timer);
          }
        }
      }, 1000);
    }
  }
};
// 初始化编辑器
const initializeEditor = () => {
  window.tinymce.init({
    target: texteditor.value,
    height: 500,
    language: "zh-Hans",
    placeholder: placeholder,
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
      "undo redo | bold italic underline | emoticons fontselect fontsizeselect forecolor backcolor | align accordion | numlist bullist outdent indent | link unlink image media table | blockquote hr removeformat | preview fullscreen", // 添加 'image' 按钮到工具栏
    menubar: false,
    branding: false,
    quickbars_image_toolbar: "rotateimage",
    // 上传图片
    images_upload_handler: (blobInfo, progress) => {
      return new Promise(async (resolve, reject) => {
        try {
          let res = await upload(blobInfo, progress);
          resolve(getPath(res.url));
        } catch (error) {
          removeFailedImage(blobInfo.blobUri());
          reject(error);
        }
      });
    },
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
  editorInstance = editor;
  editor.on("init", function () {
    show.value = false;
    editor.setContent(moduleValue);
  });
  editor.on("change", (e) => {
    emit("update:moduleValue", editor.getContent());
  });
  // editor.ui.registry.addButton("rotateimage", {
  //   text: "",
  //   icon: "rotate-left", // 使用 TinyMCE 的内置图标
  //   onAction: function () {
  //     const selectedNode = editor.selection.getNode();
  //     if (selectedNode && selectedNode.tagName === "IMG") {
  //       const currentRotation =
  //         parseInt(selectedNode.style.transform.replace(/[^0-9]/g, "")) || 0;
  //       const newRotation = (currentRotation + 90) % 360;
  //       editor.dom.setStyle(
  //         selectedNode,
  //         "transform",
  //         `rotate(${newRotation}deg)`
  //       );
  //     }
  //   },
  // });
}
// 定义一个函数来删除上传失败的图片
const removeFailedImage = (blobUri) => {
  // 使用 TinyMCE 的 API 执行一个事务以删除图片
  editorInstance.undoManager.transact(() => {
    // 使用 TinyMCE 的 DOM API 查找并删除图片
    const imgElement = editorInstance.dom.select(`img[src="${blobUri}"]`)[0];
    if (imgElement) {
      editorInstance.dom.remove(imgElement);
    }
  });
};

// 获取所有图片
const getAllImages = () => {
  const images = editorInstance.dom.select("img");
  return Array.from(images).map((img) => img.src);
};
defineExpose({
  getAllImages,
});
</script>
