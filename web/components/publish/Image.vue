<template>
  <div class="w-full rounded-xl flex flex-wrap" ref="imageList">
    <div
      :key="item.name + index"
      v-for="(item, index) in images"
      class="viewer w-32 h-32 shadow-center relative rounded-xl mr-5 mb-7"
    >
      <img
        :src="item.url"
        class="rounded-xl absolute object-cover w-full h-full"
      />
      <p class="text-sm text-center py-1 absolute -bottom-8 truncate w-full">
        {{ item.name }}
      </p>
      <div class="mask rounded-xl absolute top-0 right-0 w-full h-full">
        <Icon
          @click="openViewer(index)"
          class="cursor-pointer hover:text-[--illuminate-color]"
          size="1.5rem"
          name="ep:view"
        ></Icon>
        <Icon
          @click="remove(index)"
          class="ml-2 cursor-pointer hover:text-[--illuminate-color]"
          size="1.5rem"
          name="ep:delete"
        ></Icon>
      </div>
    </div>
    <Upload
      class="w-32 h-32 flex items-center justify-center rounded-xl shadow-center active:shadow-active"
      @uploadSuccess="uploadSuccess"
      :multiple="true"
      accept="image/*"
    >
      <div>+</div>
    </Upload>
  </div>
  <div class="flex justify-end">
    <KButton @click="toRepeat">去除重复</KButton>
    <KButton class="ml-5" @click="sort">一键排序</KButton>
  </div>
</template>

<script setup>
const { $imageViewer } = useNuxtApp();
import Sortable from "sortablejs";
const { getPath } = usePath();
const { t } = useI18n();
let images = ref([]);
let imageList = ref(null);
onMounted(() => {
  new Sortable(imageList.value, {
    animation: 150,
    ghostClass: "sortable-ghost",
    onUpdate: async (evt) => {
      const oldIndex = evt.oldIndex;
      const newIndex = evt.newIndex;
      const movedItem = images.value.splice(oldIndex, 1)[0];
      images.value.splice(newIndex, 0, movedItem);
    },
  });
});
const remove = (index) => {
  images.value.splice(index, 1);
};
const uploadSuccess = (res) => {
  images.value.push({
    name: res.name,
    url: getPath(res.url),
  });
};
const toRepeat = () => {
  let obj = {};
  images.value = images.value.reduce((cur, next) => {
    obj[next.url] ? "" : (obj[next.url] = true && cur.push(next));
    return cur;
  }, []);
};
const sort = () => {
  images.value.sort((a, b) => {
    // return a.name.localeCompare(b.name, undefined, { numeric: true });
    return a.name.localeCompare(b.name);
  });
};
// 校验参数
const checkParams = () => {
  if (!images.value.length) {
    return t("topic-create-image");
  }
  return false;
};

// 提交
const getFormData = () => {
  let list = images.value.map((item, index) => {
    return { ...item, order: index };
  });
  return {
    topicImage: {
      images: JSON.stringify(list),
    },
  };
};
function openViewer(index) {
  $imageViewer.open(images.value, index);
}
defineExpose({
  getFormData,
  checkParams,
});
</script>

<style scoped>
.mask {
  display: none;
  background: rgba(0, 0, 0, 0.5);
  transition: all 0.3s;
  align-items: center;
  justify-content: center;
}
.viewer:hover .mask {
  display: flex;
}
</style>
