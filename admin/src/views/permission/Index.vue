<template>
  <div class="h-full flex flex-col">
    <Search :keys="keys" @query="query" @reset="reset" />
    <Table
      class="mt-5 h-full"
      v-model="pages"
      :tableTitle="tableTitle"
      :tableData="tableData"
      @query="getAuthority"
      :loading="loading"
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
import { authority } from "@/api/user.js";
import Search from "@/components/Search.vue";
import Table from "@/components/Table.vue";
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
    label: "名称",
    prop: "name",
  },
  {
    label: "备注",
    prop: "remake",
  },
  {
    label: "创建时间",
    prop: "created_at",
  },
  {
    label: "更新时间",
    prop: "updated_at",
  },
];
let tableData = ref([]);
let loading = ref(false);
let pages = ref({
  page: 1,
  pageSize: 10,
  total: 0,
});
let query = (data) => {
  getAuthority(data);
};
let reset = () => {
  getAuthority();
};
async function getAuthority(query = {}) {
  loading.value = true;
  try {
    const { data } = await authority({ ...pages.value, ...query });
    console.log(data);

    tableData.value = data.list;
    pages.value.total = data.total;
  } catch (error) {
    console.log(error);
  }
  loading.value = false;
}
</script>
