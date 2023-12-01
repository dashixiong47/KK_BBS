<template>
  <el-form ref="formEl" :model="form" :rules="rules" label-width="120px">
    <el-form-item label="分组名称" prop="name">
      <el-input v-model="form.name" placeholder="请输入分组名称" />
    </el-form-item>
    <el-form-item label="分组图标">
      <Upload v-model="form.icon" />
    </el-form-item>
    <el-form-item label="排序">
      <el-input-number v-model="form.order" />
    </el-form-item>
  </el-form>
</template>

<script setup>
import Upload from "@/components/Upload.vue";
import { addGroup, updateGroup } from "@/api/group.js";
import { ref, watch } from "vue";
let props = defineProps({
  data: {
    type: Object,
    default: () => ({}),
  },
});
let editState = ref(false);
let formEl = ref(null);
let form = ref({
  name: "",
});
const rules = {
  name: [
    {
      required: true,
      message: "请输入分组名称",
      trigger: "blur",
    },
  ],
};
let submit = async () => {
  let state = await formEl.value.validate((valid) => valid);
  if (!state) throw new Error("表单验证失败");
  let fun = editState.value ? updateGroup : addGroup;
  try {
    await fun(form.value);
  } catch (error) {
    throw new Error("表单验证失败");
  }
};
let setForm = (data) => {
  form.value = data;
};
let setEditState = (state) => {
  editState.value = state;
};
defineExpose({
  submit,
  setForm,
  setEditState,
});
</script>

<style lang="scss" scoped></style>
