<template>
  <div class="w-full relative glass rounded-xl overflow-hidden">
    <DragColumn :columns="columns" :tableData="modelValue" @update="newValue => { tableData = newValue }">
      <template #coins="{ record }">
        <KInput
          class="w-20 py-1 px-4 rounded-lg text-center"
          :max="100"
          :min="0"
          v-model.number="record.coins"
        />
      </template>
      <template #index="{ index }">
        {{ index + 1 }}
      </template>
      <template #type="{ record }">
        {{ type[record.type] }}
      </template>
      <template #hidden="{ record }">
        <KSwitch v-model="record.hidden" :on-value="1" :off-value="0" />
      </template>
      <template #remake="{ record }">
        <KInput class="w-20 py-1 px-4 rounded-lg" v-model="record.remake" />
      </template>
    </DragColumn>
    <div class="w-full flex">
      <Upload
        class="w-[50%]"
        accept="*"
        multiple
        @uploadSuccess="uploadSuccess"
      >
        <KButtonNoStyle class="w-full border-r">上传</KButtonNoStyle>
      </Upload>
      <KButtonNoStyle class="w-[50%]" @click="netdiskVisible = true">
        网盘
      </KButtonNoStyle>
    </div>
  </div>
  <Dialog v-model="netdiskVisible" title="网盘" @confirm="netdiskConfirm">
    <!-- <Netdisk @select="selectNetdisk" /> -->
    <div class="w-[500px] p-5">
      <div class="mb-2">
        <div class="mb-2"><label for="1212">名称:</label></div>
        <KInput
          class="w-full p-2 rounded-lg"
          v-model="netdisk.name"
          placeholder="url"
        />
      </div>
      <div class="mb-2">
        <div class="mb-2"><label for="1212">地址:</label></div>
        <KInput
          class="w-full p-2 rounded-lg"
          v-model="netdisk.netDiskUrl"
          placeholder="url"
        />
      </div>
      <div class="mb-2">
        <div class="mb-2"><label for="1212">密码:</label></div>
        <KInput
          class="w-full p-2 rounded-lg"
          v-model="netdisk.netDiskCode"
          placeholder="url"
        />
      </div>
    </div>
  </Dialog>
</template>

<script setup>
let uploadGattachmentList = ref([]);
const props = defineProps({
  modelValue: {
    type: Array,
    default: () => [],
  },
});
let type = {
  1: "附件",
  2: "网盘",
};
// 网盘弹窗
let netdiskVisible = ref(false);
// 网盘附件
let netdisk = ref({
  netDiskUrl: null,
  netDiskCode: null,
  type: 2,
  remake: null,
});
let columns = ref([
  {
    title: "顺序",
    slot: "index",
  },
  {
    title: "名称",
    prop: "name",
    width: "200px",
  },
  {
    title: "类型",
    slot: "type",
  },
  {
    title: "评论可见",
    slot: "hidden",
  },
  {
    title: "金币",
    slot: "coins",
  },
  {
    title: "备注",
    slot: "remake",
  },
]);
const emit = defineEmits(["update:modelValue"]);
const modelValue = computed({
  get: () => {
    if (props.modelValue) {
      return props.modelValue;
    }
    return uploadGattachmentList.value;
  },
  set: (value) => {
    if (props.modelValue) {
      emit("update:modelValue", value);
      return;
    }
    uploadGattachmentList.value = value;
  },
});
const uploadSuccess = (data) => {
  modelValue.value.push({
    name: data.name,
    fileUrl: data.url,
    fileType: data.type,
    fileSize: data.size,
    coins: 0,
    hidden: 0,
    type: 1,
  });
};
const netdiskConfirm = () => {
  console.log(netdisk.value);
  modelValue.value.push({
    name: netdisk.value.name,
    netDiskUrl: netdisk.value.netDiskUrl,
    netDiskCode: netdisk.value.netDiskCode,
    coins: 0,
    hidden: 0,
    type: 2,
    remake: null,
  });
};
</script>
