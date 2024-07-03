import { defineStore } from "pinia";

export const useAdminStore = defineStore('AdminStore', {
    state: () => ({
        adminTickets: [],
    }),
    actions:{
        async getAdminTickets(token) {
            const res = await fetch("http://localhost:8080/v1/tsr/tickets",{
                method: "POST",
                body: JSON.stringify({
                    token: token,
                    mode: 'admin',
                }),
            })
            const data = await res.json()
            if (res.status === 200) {
                console.log(data.tickets)
                this.adminTickets = data.tickets
            }
        },
    }, 
})