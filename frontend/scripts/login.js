document.addEventListener('DOMContentLoaded', () => {
	// Анимация при наведении на инпуты
	document.querySelectorAll('.input-group input').forEach(input => {
		input.addEventListener('focus', () => {
			input.parentElement.querySelector('.underline').style.width = '100%'
		})

		input.addEventListener('blur', () => {
			if (!input.value) {
				input.parentElement.querySelector('.underline').style.width = '0'
			}
		})
	})

	// Валидация формы
	document.querySelector('.auth-form').addEventListener('submit', e => {
		e.preventDefault()
		const username = document.querySelector(
			'.input-group input[name="username"]'
		).value
		const password = document.querySelector(
			'.input-group input[name="password"]'
		).value

		if (!username || !password) {
			alert('Пожалуйста, заполните все поля')
			return
		}
	})
})
