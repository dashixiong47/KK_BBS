<template>
  <div class="tableHight w-full h-full">
    <el-card v-loading="props.loading">
      <template v-if="header" #header>
        <slot name="header"> </slot>
      </template>

      <el-table :data="props.tableData" style="width: 100%" border>
        <el-table-column type="index" width="50" align="center" />
        <el-table-column
          v-for="item in tableTitle"
          :prop="item.prop"
          :label="item.label"
          :width="item.width || null"
          v-bind="item.option"
        >
          <template #default="{ row, $index }" v-if="item.slot">
            <slot :name="item.slot" :row="row" :index="$index" />
          </template>
        </el-table-column>
      </el-table>
      <div class="mt-5 flex justify-end" v-if="props.modelValue">
        <el-pagination
          v-model:current-page="modelValue.page"
          v-model:page-size="modelValue.pageSize"
          layout="total, prev, pager, next, jumper"
          :total="modelValue.total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
const props = defineProps({
  tableTitle: {
    type: Array,
    default: () => [],
  },
  tableData: {
    type: Array,
    default: () => [],
  },
  modelValue: {
    type: Object,
    default: () => {},
  },
  header: {
    type: Boolean,
    default: false,
  },
  loading: {
    type: Boolean,
    default: false,
  },
});
let handleSizeChange = (val) => {
  emit("update:modelValue", {
    ...props.modelValue,
    pageSize: val,
  });
  emit("query");
};
let handleCurrentChange = (val) => {
  emit("update:modelValue", {
    ...props.modelValue,
    page: val,
  });
  emit("query");
};
let emit = defineEmits(["update:modelValue", "query"]);
onMounted(() => {
  let content = document.querySelector(".tableHight");
  emit("update:modelValue", {
    ...props.modelValue,
    pageSize: Number((content.offsetHeight / 40).toFixed(0)),
  });
  emit("query");
});
</script>
