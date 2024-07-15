import { defineStore } from "pinia";
import settings from '../settings.json'
const path = settings.url

export const useArchiveStore = defineStore('ArchiveStore', {
        state: () => ({
                archiveTickets: [],
                filteredTickets: [],
            }),
            actions: {
                async getArchiveTickets(token) {
                    const res = await fetch(path+'tsr/tickets',{
                        method: "POST",
                        body: JSON.stringify({
                            token: token,
                            mode: 'archive',
                        }),
                    })
                    const data = await res.json()
                    if (res.status === 200) {
                        console.log(data.tickets)
                        this.archiveTickets = data.tickets
                    }
                },
            }

})