import type { AdminGroup, GroupPlatform, ProviderBrand } from '@/types'

export const providerBrands: Array<{ value: ProviderBrand; label: string }> = [
  { value: 'deepseek', label: 'DeepSeek' },
  { value: 'qwen', label: 'Qwen' },
  { value: 'minimax', label: 'MiniMax' },
  { value: 'seedance', label: 'Seedance' },
  { value: 'mimo', label: 'MiMo' },
  { value: 'glm', label: 'GLM' },
  { value: 'happyhorse', label: 'HappyHorse' },
]

const providerBrandSet = new Set<string>(providerBrands.map(({ value }) => value))

export const isProviderBrand = (value: unknown): value is ProviderBrand =>
  typeof value === 'string' && providerBrandSet.has(value)

export const providerBrandLabel = (brand: ProviderBrand): string =>
  providerBrands.find(({ value }) => value === brand)?.label ?? brand

export const groupProviderBrand = (
  group: Pick<AdminGroup, 'models_list_config'> & Partial<Pick<AdminGroup, 'name'>>,
): ProviderBrand | undefined => {
  const brand = group.models_list_config?.provider_brand
  if (isProviderBrand(brand)) return brand

  const name = group.name?.toLowerCase().replace(/[\s_-]/g, '') ?? ''
  return providerBrands.find(({ value }) => name.includes(value))?.value
}

export const providerTransportPlatform = (_brand: ProviderBrand): GroupPlatform => 'openai'
