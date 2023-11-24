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
          :key="rowIndex+new Date().getTime()" 
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
import Sortable from 'sortablejs';

const props = defineProps({
  columns: { type: Array, default: () => [] },
  tableData: { type: Array, default: () => [] },
  width: { type: String, default: '100%' },
  height: { type: String, default: 'auto' },
});

const tableRef = ref(null);

const emit = defineEmits(['update']);

onMounted(() => {
  new Sortable(tableRef.value, {
    animation: 150,
    ghostClass: 'sortable-ghost',
    onUpdate: async (evt) => {
      const oldIndex = evt.oldIndex;
      const newIndex = evt.newIndex;
      let updatedTableData = [...props.tableData];
      const movedItem = updatedTableData.splice(oldIndex, 1)[0];
      updatedTableData.splice(newIndex, 0, movedItem);
      emit('update', updatedTableData);
    },
  });
});
</script>

<style scoped>
.sortable-ghost {
  opacity: 0.5;
}
</style>
