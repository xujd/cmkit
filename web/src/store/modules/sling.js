import { queryDict } from '@/api/sys'

const state = {
  slingTypes: [],
  slingTons: [],
  slingUseStatus: [],
  slingInspectStatus: []
}

const mutations = {
  SET_DICT: (state, data) => {
    state.slingTypes = data.slingTypes
    state.slingTons = data.slingTons
    state.slingUseStatus = data.slingUseStatus
    state.slingInspectStatus = data.slingInspectStatus
  }
}

const getters = {
  slingTypes: state => state.slingTypes,
  slingTons: state => state.slingTons,
  slingUseStatus: state => state.slingUseStatus,
  slingInspectStatus: state => state.slingInspectStatus
}

const actions = {
  initDict({ commit }) {
    return new Promise((resolve, reject) => {
      queryDict('Sling', '').then(response => {
        const { data } = response
        const result = {}
        result.slingTypes = data.filter(item => item.type === 'SLING_TYPE')
        result.slingTons = data.filter(item => item.type === 'TON_TYPE')
        result.slingUseStatus = data.filter(item => item.type === 'USE_STATUS_TYPE')
        result.slingInspectStatus = data.filter(item => item.type === 'INSPECT_STATUS_TYPE')
        commit('SET_DICT', result)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  getters,
  actions
}
