<template>
  <el-form :model="data">
    <el-form-item
      :label="item.label"
      :label-width="width + 'rem'"
      v-for="item in props.col"
      :key="item.label"
      style="width: 24rem"
      :prop="item.prop"
    >
      <el-text v-if="item.type == 'label'" />
      <el-date-picker
        v-model="data[item.prop]"
        type="date"
        placeholder="选择日期"
        v-if="item.type == 'date'"
      />
      <el-input v-model="data[item.prop]" v-if="item.type == 'string'" />
      <el-input
        v-model.number="data[item.prop]"
        type="number"
        v-if="item.type == 'number'"
      />
      <el-radio-group v-model="data[item.prop]" v-if="item.type == 'select'">
        <el-radio
          v-for="option in item.options"
          :key="option.label"
          :label="option.value"
          >{{ option.label }}</el-radio
        >
      </el-radio-group>
      <el-select
        v-model="data[item.prop]"
        multiple
        placeholder="选择"
        style="width: 240px"
        v-if="item.type == 'multi-select'"
      >
        <el-option
          v-for="option in item.options"
          :key="option.value"
          :label="option.label"
          :value="option.value"
        />
      </el-select>
      <el-select
        v-model="data[item.prop]"
        placeholder="选择"
        v-if="item.type == 'single-select-cascader'"
      >
        <el-option-group
          v-for="group in item.options"
          :key="group.label"
          :label="group.label"
        >
          <el-option
            v-for="item in group.options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          ></el-option
        ></el-option-group>
      </el-select>
      <el-select
        v-model="data[item.prop]"
        placeholder="选择"
        style="width: 240px"
        v-if="item.type == 'single-select'"
      >
        <el-option
          v-for="option in item.options"
          :key="option.value"
          :label="option.label"
          :value="option.value"
        />
      </el-select>
    </el-form-item>
  </el-form>
  <span class="dialog-footer">
    <el-button @click="() => emit('submit', data)" type="primary"
      >确定</el-button
    >
  </span>
</template>

<script lang="ts" setup>
import { reactive } from "vue";

let props = defineProps(["col"]);
let emit = defineEmits(["submit"]);
let data = reactive(<any>{});
let width = ref(6);

onMounted(() => {
  props.col.forEach((item: any) => {
    if (item.label.length > width.value) {
      width.value = item.label.length + 1;
    }
  });
});
</script>
