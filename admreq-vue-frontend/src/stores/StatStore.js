import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useStatStore = defineStore('StatStore', {
        state: () => ({
                statData: [],
            }),
            actions: {
                async getStatData(target, token) {
                    const res = await fetch(path+'tsr/stat',{
                        method: "POST",
                        body: JSON.stringify({
                            target_dep: target,
                            token: token,
                        }),
                    })
                    const data = await res.json()
                    if (res.status === 200) {
                        this.statData = data.byDepartment
                        console.log(this.statData)
                    }
                },
            }

})