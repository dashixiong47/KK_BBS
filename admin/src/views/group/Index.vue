<template>
  <div class="flex flex-col">
    <Search :keys="keys" @query="query" @reset="reset" />
    <Table
      class="mt-5"
      v-model="pages"
      :tableTitle="tableTitle"
      :tableData="tableData"
      @query="getGroupData"
      :header="true"
    >
      <template #header>
        <el-button type="primary" @click="add">添加</el-button>
      </template>
      <template #icon="{ row }">
        <span v-if="!row.icon">--</span>
        <img
          v-else
          class="h-8 w-8 rounded-full block m-auto"
          :src="getPath(row.icon)"
          alt=""
          srcset=""
        />
      </template>
      <template #updated_at="{ row }">
        {{ formatTime(row.updated_at) }}
      </template>
      <template #option="{ row }">
        <el-button type="text" @click="edit(row)">编辑</el-button>
        <el-button type="text" @click="_delete(row)">删除</el-button>
      </template>
    </Table>
  </div>
  <el-drawer v-model="drawer" :size="500">
    <Form ref="form" :data="data" />
    <template #footer>
      <div style="flex: auto">
        <el-button @click="cancelClick">取消</el-button>
        <el-button type="primary" @click="confirmClick">提交</el-button>
      </div>
    </template>
  </el-drawer>
</template>

<script setup>
import { ref } from "vue";
import { group, deleteGroup } from "@/api/group.js";
import Search from "@/components/Search.vue";
import Table from "@/components/Table.vue";
import { getPath } from "@/utils/path.js";
import { formatTime } from "@/utils/time.js";
import Form from "./form.vue";
let drawer = ref(false);
let form = ref(null);
let data = ref({});
let keys = [
  {
    label: "名称",
    type: "input",
    placeholder: "请输入名称",
    key: "name",
    value: "",
  },
];
let tableTitle = [
  {
    label: "图标",
    prop: "icon",
    slot: "icon",
    width: "80",
    option: {
      align: "center",
    },
  },
  {
    label: "名称",
    prop: "name",
  },
  {
    label: "排序",
    prop: "order",
  },
  {
    label: "创建时间",
    prop: "updated_at",
    slot: "updated_at",
  },
  {
    label: "编辑",
    prop: "option",
    slot: "option",
  },
];
let tableData = ref([]);
let pages = ref({
  page: 1,
  pageSize: 10,
  total: 0,
});
function add() {
  data.value = {};
  drawer.value = true;
  form.value.setForm({});
  form.value.setEditState(false);
}
function edit(row) {
  data.value = row;
  drawer.value = true;
  setTimeout(() => {
    form.value.setForm(row);
    form.value.setEditState(true);
  }, 0);
}
async function _delete(row) {
  try {
    await deleteGroup({ id: row.id });
    getGroupData();
  } catch (error) {
    console.log(error);
  }
}
function cancelClick() {
  drawer.value = false;
}
async function confirmClick() {
  try {
    let state = await form.value.submit();
    getGroupData();
  } catch (error) {
    console.log(error);
  }
  drawer.value = false;
}
let query = (data) => {
  getGroupData(data);
};
let reset = () => {
  getGroupData();
};
async function getGroupData(query = {}) {
  try {
    const { data } = await group({ ...pages.value, ...query });
    tableData.value = data.list.sort((a, b) => a.order - b.order);
    pages.value.total = data.total;
  } catch (error) {}
}
</script>
