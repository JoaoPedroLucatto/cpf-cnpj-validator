<template>
  <div
    v-if="message"
    :class="['alert', typeClass]"
    role="alert"
  >
    {{ message }}
  </div>
</template>

<script setup lang="ts">
import { computed, defineProps } from "vue";

const props = defineProps({
  message: {
    type: String,
    required: false,
    default: "",
  },
  type: {
    type: String,
    default: "success",
  },
});

const typeClass = computed(() =>
  ({
    success: "alert-success",
    danger: "alert-danger",
    warning: "alert-warning",
    info: "alert-info",
  }[props.type])
);
</script>
