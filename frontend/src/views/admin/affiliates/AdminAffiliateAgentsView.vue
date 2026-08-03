<template>
  <AppLayout>
    <div class="space-y-4">
      <div class="flex flex-wrap items-center gap-3">
        <input
          v-model="search"
          class="input w-full md:w-80"
          :placeholder="t('admin.affiliates.agents.searchPlaceholder')"
          @input="debounceLoad"
        >
        <button class="btn btn-secondary" :disabled="loading" @click="load">
          {{ t('common.refresh') }}
        </button>
      </div>

      <div class="card overflow-x-auto">
        <table class="w-full min-w-[720px] text-left text-sm">
          <thead>
            <tr class="border-b border-gray-200 text-gray-500 dark:border-dark-700 dark:text-dark-400">
              <th class="px-4 py-3 font-medium">{{ t('admin.affiliates.agents.agent') }}</th>
              <th class="px-4 py-3 text-right font-medium">{{ t('admin.affiliates.agents.inviteeCount') }}</th>
              <th class="px-4 py-3 text-right font-medium">{{ t('admin.affiliates.agents.weekSales') }}</th>
              <th class="px-4 py-3 text-right font-medium">{{ t('admin.affiliates.agents.monthSales') }}</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="loading">
              <td colspan="4" class="px-4 py-10 text-center text-gray-500">{{ t('common.loading') }}</td>
            </tr>
            <tr v-else-if="items.length === 0">
              <td colspan="4" class="px-4 py-10 text-center text-gray-500">{{ t('admin.affiliates.agents.empty') }}</td>
            </tr>
            <template v-else>
              <tr
                v-for="item in items"
                :key="item.agent_id"
                class="border-b border-gray-100 last:border-0 dark:border-dark-800"
              >
                <td class="px-4 py-3">
                  <div class="font-medium text-gray-900 dark:text-white">{{ item.email }}</div>
                  <div class="text-gray-500 dark:text-dark-400">{{ item.username || '-' }}</div>
                </td>
                <td class="px-4 py-3 text-right">{{ item.invitee_count.toLocaleString() }}</td>
                <td class="px-4 py-3 text-right font-medium text-emerald-600 dark:text-emerald-400">
                  {{ formatCurrency(item.week_sales) }}
                </td>
                <td class="px-4 py-3 text-right font-semibold text-primary-600 dark:text-primary-400">
                  {{ formatCurrency(item.month_sales) }}
                </td>
              </tr>
            </template>
          </tbody>
        </table>
      </div>

      <Pagination
        v-if="total > 0"
        :page="page"
        :total="total"
        :page-size="pageSize"
        @update:page="changePage"
        @update:pageSize="changePageSize"
      />
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import AppLayout from '@/components/layout/AppLayout.vue'
import Pagination from '@/components/common/Pagination.vue'
import { affiliatesAPI, type AffiliateAgentSales } from '@/api/admin/affiliates'
import { useAppStore } from '@/stores/app'
import { formatCurrency } from '@/utils/format'

const { t } = useI18n()
const appStore = useAppStore()
const items = ref<AffiliateAgentSales[]>([])
const loading = ref(false)
const search = ref('')
const page = ref(1)
const pageSize = ref(20)
const total = ref(0)
let timer: ReturnType<typeof setTimeout> | undefined

async function load() {
  loading.value = true
  try {
    const result = await affiliatesAPI.listAgentSales({
      page: page.value,
      page_size: pageSize.value,
      search: search.value.trim(),
    })
    items.value = result.items || []
    total.value = result.total || 0
  } catch (error: any) {
    appStore.showError(error?.message || t('admin.affiliates.errors.loadFailed'))
  } finally {
    loading.value = false
  }
}

function debounceLoad() {
  clearTimeout(timer)
  timer = setTimeout(() => {
    page.value = 1
    void load()
  }, 300)
}

function changePage(value: number) {
  page.value = value
  void load()
}

function changePageSize(value: number) {
  pageSize.value = value
  page.value = 1
  void load()
}

onMounted(() => void load())
</script>
