<template>
  <div id="app">
    <b-navbar type="dark" variant="dark">
      <b-container>
        <b-navbar-brand to="/">
          <img src="@/assets/SpyServer.png" width="30" />
        </b-navbar-brand>

        <b-navbar-nav>
          <b-nav-item to="/">New domain</b-nav-item>
          <b-nav-item to="/previous">Previous domains</b-nav-item>
        </b-navbar-nav>    
      </b-container>
    </b-navbar>
    <router-view/>
  </div>
</template>

<script>
export default {
  
}
</script>

<style>
</style>

