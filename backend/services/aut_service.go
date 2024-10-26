package services

import (
	"clientapp/models"
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService es responsable de manejar la lógica de negocio relacionada con la autenticación
type AuthService struct {
	DB *gorm.DB
}

// Regiter registra un nuevo usuario en la base de datos
// Se valida que el correo sea unico, que las contraseñas coincidan y que se cumpla con los requisitos
func (s *AuthService) Register(user *models.User, passwordConfirm string) error {

	// Verificamos si la contraseña cumple con los criterios de seguridad
	if err := validatePassword(user.Password); err != nil {
		return err // Devolvemos un error si no es valida
	}

	// Comprobamos si las contraseñas coinciden
	if user.Password != passwordConfirm {
		return errors.New("las contraseñas no coinciden")
	}

	// Verificamos que el correo no esté ya registrado
	var existingUser models.User
	if err := s.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("el correo ya está registrado")
	}

	// Hasheamos la contraseña antes de almacenarla en la base de datos
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("ha ocurrido un error, reintentelo de nuevo")
	}

	// Asignamos la password hasheada al usuario
	user.Password = string(hashedPassword)

	// Guardamos el nuevo usuario en la base de datos
	if err := s.DB.Create(user).Error; err != nil {
		return errors.New("ha ocurrido un error al guardar, reintentelo de nuevo")
	}

	return nil // Si todo salió bien, devolvemos nil (sin errores)
}

// Login permite que un usuario inicie sesion
// Verifica que el correo y la contraseña coinciden en la base de datos
func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User

	// Buscamos al usuario por su correo electronico
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("el usuario no existe")
	}

	// Comparamos la contraseña con la contraseña hasheada almacenada
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("credenciales invalidas")
	}

	return &user, nil // Si la autenticacion es exitosa, devolvemos los datos del usuario

}

// validatePassword verifica que la contraseña cumple con los requisitos de seguridad
func validatePassword(password string) error {

	// La contraseña debe tener al menos 8 caracteres
	if len(password) < 8 {
		return errors.New("la contraseña debe tener al menos 8 caracteres")
	}

	var hasUpper bool
	var hasNumber bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if !hasUpper {
		return errors.New("la contraseña debe tener al menos una mayuscula")
	}
	if !hasNumber {
		return errors.New("la contraseña debe tener al menos un numero")
	}

	return nil // SI la contraseña es valida no devolvemos errores
}
