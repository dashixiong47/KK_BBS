<template>
  <el-card class="flex-shrink-0">
    <el-form :inline="true" :model="formData" class="demo-form-inline">
      <el-form-item :label="item.label" v-for="item in keys">
        <component
          :is="type[item.type]"
          v-model="formData[item.key]"
          :placeholder="item.placeholder"
          clearable
        />
      </el-form-item>
      <div class="flex justify-end">
        <el-button type="primary" @click="query">{{
          $t("search-query")
        }}</el-button>
        <el-button @click="reset">{{ $t("search-reset") }}</el-button>
      </div>
    </el-form>
  </el-card>
</template>

<script setup>
import { ref } from "vue";
import { ElInput, ElButton } from "element-plus";
import "element-plus/dist/index.css";
// import { t } from "vue-i18n";
const props = defineProps({
  /**
   * 搜索条件
   * @type {Array}
   * @default []
   * @example
   * [
   *  {
   *   label: "姓名",
   *   type: "input",
   *   placeholder: "请输入姓名",
   *   key: "name",
   *   value: "",
   * }
   * ]
   */
  keys: {
    type: Array,
    default: () => [],
  },
});
let formData = ref({});
let type = {
  input: ElInput,
};
let emit = defineEmits(["query", "reset"]);
function query() {
  emit("query", formData.value);
}
function reset() {
  formData.value = {};
  emit("reset");
}
</script>
