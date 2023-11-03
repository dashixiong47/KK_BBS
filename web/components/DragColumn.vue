<template>
  <div :style="{ width: width, height: height }" class="overflow-x-auto">
    <table class="w-full">
      <colgroup>
        <col
          v-for="(column, index) in columns"
          :key="index"
          :style="{ width: column.width, minWidth: column.minWidth, maxWidth: column.maxWidth }"
        />
      </colgroup>
      <thead class="border-b">
        <tr>
          <th v-for="(column, index) in columns" :key="index" class="p-4 whitespace-nowrap">
            {{ column.title }}
          </th>
        </tr>
      </thead>
      <tbody ref="tableRef">
        <tr
          v-for="(record, rowIndex) in tableData"
          :key="rowIndex"
          class="relative w-full border-b"
        >
          <td
            v-for="(column, colIndex) in columns"
            :key="colIndex"
            class="text-center p-4 overflow-hidden"
          >
            <div class="truncate" :style="{ width: columns[colIndex].width, minWidth: columns[colIndex].minWidth, maxWidth: columns[colIndex].maxWidth }" v-if="!column.slot">{{ record[column.prop] }}</div>
            <slot
              v-else
              :name="column.slot"
              :record="record"
              :index="rowIndex"
              :column="column"
            >
            </slot>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="!tableData.length" class="w-full text-center border-b py-4">
      No Data
    </div>
  </div>
</template>

<script setup>
import Sortable from "sortablejs";

const { columns, tableData, width, height } = defineProps({
  columns: { type: Array, default: () => [] },
  tableData: { type: Array, default: () => [] },
  width: { type: String, default: "100%" },
  height: { type: String, default: "auto" },
});

const tableRef = ref(null);
const emit = defineEmits(["update:modelValue"]);
onMounted(() => {
  new Sortable(tableRef.value, {
    animation: 150,
    ghostClass: "sortable-ghost",
    direction: "vertical",  // 仅允许垂直拖动
    onUpdate(evt) {
      const oldIndex = evt.oldIndex;
      const newIndex = evt.newIndex;
      let tableDatas = [...tableData];
      const movedItem = tableDatas.splice(oldIndex, 1)[0];
      tableDatas.splice(newIndex, 0, movedItem);
      emit("update:modelValue", tableDatas);
    },
  });
});
</script>

<style scoped>
/* Tailwind CSS does not have built-in classes for the "sortable-ghost" class used by Sortable.js. 
   You might need to customize this class based on your requirements. */
.sortable-ghost {
  opacity: 0.5;
}
</style>
