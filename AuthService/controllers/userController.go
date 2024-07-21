package controllers

import (
        "alumniportal.com/shared/initializers"
        "alumniportal.com/shared/models"
        "github.com/gin-gonic/gin"
        "github.com/golang-jwt/jwt"
        "github.com/sirupsen/logrus"
        "golang.org/x/crypto/bcrypt"
        "net/http"
        "os"
        "time"
)

var body struct {
        Name     string `json:"name"`
        LastName string `json:"lastName"`
        Email    string `json:"email" binding:"required"`
        Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
        if c.ShouldBindJSON(&body) != nil {
                logrus.WithFields(logrus.Fields{
                        "error": "Failed to read body",
                }).Error("Signup error")
                c.JSON(http.StatusBadRequest, gin.H{
                        "error": "Failed to read body.",
                })
                return
        }

        hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
        if err != nil {
                logrus.WithFields(logrus.Fields{
                        "error": err.Error(),
                }).Error("Failed to hash password")
                c.JSON(http.StatusInternalServerError, gin.H{
                        "error": "Failed to hash password",
                })
                return
        }

        user := models.User{Name: body.Name, LastName: body.LastName, Email: body.Email, Password: string(hash)}

        result := initializers.DB.Create(&user)

        if result.Error != nil {
                logrus.WithFields(logrus.Fields{
                        "error": result.Error.Error(),
                }).Error("Failed to create user")
                c.JSON(http.StatusInternalServerError, gin.H{
                        "error": "Failed to create user",
                })
                return
        }

        logrus.WithFields(logrus.Fields{
                "user_id": user.ID,
        }).Info("User created successfully")
        c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
        if c.Bind(&body) != nil {
                logrus.WithFields(logrus.Fields{
                        "error": "Failed to read body",
                }).Error("Login error")
                c.JSON(http.StatusBadRequest, gin.H{
                        "error": "Failed to read body.",
                })
                return
        }

        var user models.User
        initializers.DB.First(&user, "email = ?", body.Email)

        if user.ID == 0 {
                logrus.WithFields(logrus.Fields{
                        "email": body.Email,
                }).Warn("Invalid email or password")
                c.JSON(http.StatusUnauthorized, gin.H{
                        "error": "Invalid Email or Password",
                })
                return
        }

        err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

        if err != nil {
                logrus.WithFields(logrus.Fields{
                        "email": body.Email,
                }).Warn("Invalid email or password")
                c.JSON(http.StatusUnauthorized, gin.H{
                        "error": "Invalid Email or Password",
                })
                return
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
                "sub": user.ID,
                "exp": time.Now().Add(time.Hour * 48).Unix(),
        })

        tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

        if err != nil {
                logrus.WithFields(logrus.Fields{
                        "error": err.Error(),
                }).Error("Failed to sign token")
                c.JSON(http.StatusBadRequest, gin.H{
                        "error": err.Error(),
                })
                return
        }

        c.SetSameSite(http.SameSiteNoneMode)
        c.SetCookie("Authorization", tokenString, 3600*24, "/", "alumni-portal.ru", true, true)
        logrus.WithFields(logrus.Fields{
                "user_id": user.ID,
        }).Info("User logged in successfully")

        var role string

        if user.IsAdmin && user.Verified {
                role = "Admin"
        } else {
                if user.IsAlumni && user.Verified {
                        role = "Alumni"
                } else {
                        if user.Verified {
                                role = "Student"
                        } else {
                                role = "Unverified"
                        }
                }
        }

        c.JSON(http.StatusOK, gin.H{
                "userId":                     user.ID,
                "name":                       user.Name,
                "lastName":                   user.LastName,
                "email":                      user.Email,
                "specialization":             user.Specialization,
                "availableCustomerInterview": user.AvailableCustdev,
                "description":                user.Description,
                "portfolioLink":              user.PortfolioLink,
                "socialsLink":                user.SocialsLink,
                "role":                       role,
        })

}

func Validate(c *gin.Context) {
        user, _ := c.Get("user")

        if user == nil {
                logrus.Warn("User not found during validation")
                c.JSON(http.StatusUnauthorized, gin.H{
                        "error": "User not found",
                })
                return
        }

        logrus.WithFields(logrus.Fields{
                "user_id": user.(models.User).ID,
        }).Info("User validated successfully")
        c.JSON(http.StatusOK, gin.H{
                "message": user,
        })
}

func GetInfo(c *gin.Context) {
        user, _ := c.Get("user")

        if user == nil {
                logrus.Warn("User not found during info retrieval")
                c.JSON(http.StatusUnauthorized, gin.H{
                        "error": "User not found",
                })
                return
        }

        logrus.WithFields(logrus.Fields{
                "user_id": user.(models.User).ID,
        }).Info("User info retrieved successfully")

        var role string
        
        if user.(models.User).IsAdmin && user.(models.User).Verified {
                role = "Admin"
        } else {
                if user.(models.User).IsAlumni && user.(models.User).Verified {
                        role = "Alumni"
                } else {
                        if user.(models.User).Verified {
                                role = "Student"
                        } else {
                                role = "Unverified"
                        }
                }
        }
        
        c.JSON(http.StatusOK, gin.H{
                "userId":                         user.(models.User).ID,
                "name":                           user.(models.User).Name,
                "lastName":                       user.(models.User).LastName,
                "email":                          user.(models.User).Email,
                "role":                           role,
                "specialization":                 user.(models.User).Specialization,
                "availableCustomerInterview":     user.(models.User).AvailableCustdev,
                "portfolioLink":                  user.(models.User).PortfolioLink,
                "socialsLink":                    user.(models.User).SocialsLink,
                "isAlumni":                       user.(models.User).IsAlumni,
                "isAdmin":                        user.(models.User).IsAdmin,
        })
}
