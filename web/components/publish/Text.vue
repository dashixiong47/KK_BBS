<template>
  <div class="shadow-center rounded-xl overflow-hidden">
    <DragColumn
      :columns="columns"
      :tableData="tableData"
      @update="
        (newValue) => {
          tableData = newValue;
        }
      "
    >
      <template #index="{ index }">
        {{ index + 1 }}
      </template>
      <template #img="{ record }">
        <div class="flex justify-center">
          <img class="h-20" :src="record.url" alt="" srcset="" />
        </div>
      </template>
    </DragColumn>
    <div class="w-full flex">
      <Upload
        class="w-full"
        accept=".txt"
        multiple
        @uploadSuccess="uploadSuccess"
      >
        <KButtonNoStyle class="w-full border-r">上传</KButtonNoStyle>
      </Upload>
    </div>
    <div @click="toRepeat">去除重复</div>
    <div @click="sort">一键排序</div>
  </div>
</template>

<script setup>
const { getPath } = usePath();
const {getSize} = useSize();
let columns = ref([
  {
    title: "顺序",
    slot: "index",
    width: "100px",
  },
  {
    title: "名称",
    prop: "name",
    width: "200px",
  },
  {
    title: "大小",
    prop: "size",
  },
]);
let tableData = ref([]);
const uploadSuccess = (res) => {
  tableData.value.push({
    name: res.name,
    url: getPath(res.url),
    size: getSize(res.size)+"M",
  });
};
const toRepeat = () => {
  let obj = {};
  tableData.value = tableData.value.reduce((cur, next) => {
    obj[next.url] ? "" : (obj[next.url] = true && cur.push(next));
    return cur;
  }, []);
};
const sort = () => {
  tableData.value.sort((a, b) => {
    return a.name.localeCompare(b.name);
  });
};
// 校验参数
const checkParams = () => {
  if (!formData.value.topicBasic.content) {
    return t("topic-create-content");
  }
  return false;
};

// 提交
const getFormData = () => {
  let list = tableData.value.map((item, index) => {
    return { ...item, order: index };
  });
  return {
    topicText: {
        texts: JSON.stringify(list),
    },
  };
};
defineExpose({
  getFormData,
  checkParams,
});
</script>
