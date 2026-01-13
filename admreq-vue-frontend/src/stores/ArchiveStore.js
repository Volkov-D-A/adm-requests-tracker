import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useArchiveStore = defineStore('ArchiveStore', {
        state: () => ({
                archiveTickets: [],
                archivePages: 0,
            }),
            actions: {
                async getArchiveTickets(token, page) {
                    const res = await fetch(path+'tsr/archiv',{
                        method: "POST",
                        body: JSON.stringify({
                            token: token,
                            page: page,
                        }),
                    })
                    const data = await res.json()
                    if (res.status === 200) {
                        this.archiveTickets = data.tickets
                        this.archivePages = data.pages
                    }
                },
            }

})