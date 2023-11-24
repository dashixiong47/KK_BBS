<template>
  <div class="h-full flex flex-col">
    <Search :keys="keys" @query="query" @reset="reset" />
    <Table
      class="mt-5 h-full"
      v-model="pages"
      :tableTitle="tableTitle"
      :tableData="tableData"
      @query="getUser"
    >
      <template #avatar="{ row }">
        <span v-if="!row.avatar">--</span>
        <img
          v-else
          class="h-8 w-8 rounded-full block m-auto"
          :src="row.avatar"
          alt=""
          srcset=""
        />
      </template>
    </Table>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { users } from "@/api/user.js";
import Search from "@/components/Search.vue";
import Table from "@/components/Table.vue";
let keys = [
  {
    label: "姓名",
    type: "input",
    placeholder: "请输入姓名",
    key: "name",
    value: "",
  },
];
let tableTitle = [
  {
    label: "头像",
    prop: "avatar",
    slot: "avatar",
    width: "80",
    option: {
      align: "center",
    },
  },
  {
    label: "名称",
    prop: "nickname",
  },
  {
    label: "账号",
    prop: "username",
  },
  {
    label: "邮箱",
    prop: "email",
  },
  {
    label: "权限",
    prop: "role",
  },
];
let tableData = ref([
  {
    name: "王小虎",
    age: 18,
    address: "上海市普陀区金沙江路 1518 弄",
  },
  {
    name: "王小虎",
    age: 18,
    address: "上海市普陀区金沙江路 1518 弄",
  },
  {
    name: "王小虎",
    age: 18,
    address: "上海市普陀区金沙江路 1518 弄",
  },
]);
let pages = ref({
  page: 1,
  pageSize: 10,
  total: 0,
});
let query = (data) => {
  console.log(data);
};
let reset = () => {
  console.log("reset");
};
async function getUser() {
  try {
    const { data } = await users(pages.value);
    console.log(data);

    tableData.value = data.list;
    pages.value.total = data.total;
  } catch (error) {}
}
</script>
