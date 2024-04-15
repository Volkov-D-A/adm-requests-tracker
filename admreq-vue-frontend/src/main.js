//pinia
import { createPinia } from 'pinia'

//Vuetify
import '@mdi/font/css/materialdesignicons.css' // Ensure you are using css-loader
import 'vuetify/styles'
import { createVuetify} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
    icons: {
        defaultSet: 'mdi', // This is already the default value - only for display purposes
    },
    components,
    directives,
})

//Router
import router from './router'

import { createApp } from 'vue'
import App from './App.vue'

createApp(App).use(createPinia()).use(router).use(vuetify).mount('#app')
