<template>
	<MyHeader/>
	<BodyComponent/>
	<MyFooter/>
</template>

<script>
import MyHeader from "./components/MyHeader.vue"
import BodyComponent from "./components/BodyComponent.vue"
import MyFooter from "./components/MyFooter.vue"

export default {
	name: 'App',
	components: {
		MyHeader,
		BodyComponent,
		MyFooter,
	}
}
</script>

<style>
</style>
