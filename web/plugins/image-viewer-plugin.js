// plugins/image-viewer-plugin.js
import { defineNuxtPlugin } from '#app';
import ImageViewer from '~/components/ImageViewer.vue';
import { h, createApp, ref } from 'vue';

export default defineNuxtPlugin(() => {
  let imageViewerInstance = null;

  const openImageViewer = (images,index) => {
    if (!imageViewerInstance) {
      const imageViewerApp = createApp(ImageViewer, { images });
      const imageViewerElement = document.createElement('div');
      document.body.appendChild(imageViewerElement);
      imageViewerInstance = imageViewerApp.mount(imageViewerElement);

      // 确保在组件实例可用后调用 open 方法
      imageViewerInstance.$nextTick(() => {
        imageViewerInstance.open(images,index);
      });
    } else {
      // 如果实例已存在，则直接调用 open
      imageViewerInstance.open(images,index);
    }
  };

  const closeImageViewer = () => {
    if (imageViewerInstance) {
      imageViewerInstance.close();
      imageViewerInstance.$el.parentNode.removeChild(imageViewerInstance.$el);
      imageViewerInstance = null;
    }
  };

  return {
    provide: {
      imageViewer: {
        open: openImageViewer,
        close: closeImageViewer
      }
    }
  };
});
