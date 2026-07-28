import { describe, expect, it } from 'vitest'
import { groupProviderBrand, providerTransportPlatform } from '@/utils/providerBrands'

describe('providerBrands', () => {
  it('keeps provider identity separate from its OpenAI-compatible transport', () => {
    expect(groupProviderBrand({
      name: 'DeepSeek 官方',
      models_list_config: { enabled: true, models: [], provider_brand: 'deepseek' },
    })).toBe('deepseek')
    expect(providerTransportPlatform('deepseek')).toBe('openai')
  })

  it('recognizes legacy groups by name', () => {
    expect(groupProviderBrand({
      name: 'Happy Horse 官方',
      models_list_config: { enabled: false, models: [] },
    })).toBe('happyhorse')
  })
})
