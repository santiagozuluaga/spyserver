<template>
    <b-container class="previous">
        <b-row>
            <b-col>
                <b-card header="Dominios" header-text-variant="light" header-bg-variant="dark" header-class="card-title">
                    <b-table class="text-center" striped :items="domains"></b-table>
                </b-card>
            </b-col>
        </b-row>
    </b-container>
</template>

<script>
import axios from "axios"

export default {
    data () {
        return {
            domains: []
        }
    },
    created () {

        axios.get("api/servers/previous")
        .then(response => {
            (response.data).forEach(domain => {
            
                this.domains.push({
                    "domain": domain.domain,
                    "servers": domain.info.servers.length,
                    "ssl_grade": domain.info.ssl_grade,
                    "previous_ssl_grade": domain.info.previous_ssl_grade,
                    "updated": domain.updated
                })
            })

        })
        .catch(err => console.error(err))
    },
    methods: {

    }
}
</script>

<style>
.previous{
    margin-top: 40px;
}
</style>

<!--
<table class="table table-striped table-bordered">
    <thead>
        <tr>
            <th>Domain</th>
            <th>Servers</th>
            <th>Ssl grade</th>
            <th>Previous ssl grade</th>
            <th>Updated</th>
        </tr>
    </thead>
    <tbody v-if="domains.length == 0">
        <tr>
            <td>Vacio</td>
            <td>Vacio</td>
            <td>Vacio</td>
            <td>Vacio</td>
            <td>Vacio</td>
        </tr>
    </tbody>
    <tbody v-if="domains.length > 0">
        <tr v-bind:key="domain.domain" v-for="domain in domains">
            <td>{{domain.domain}}</td>
            <td>{{domain.servers}}</td>
            <td>{{domain.ssl_grade}}</td>
            <td>{{domain.previous_ssl_grade}}</td>
            <td>{{domain.updated}}</td>
        </tr>
    </tbody>
</table>
-->