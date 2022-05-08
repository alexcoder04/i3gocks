# Maintainer: alexcoder04 <alexcoder04@protonmail.com>
pkgname=kherson-git
_pkgname=kherson
pkgver=1.0.0
pkgrel=1
epoch=
pkgdesc="minimal status line generator for i3/sway written in go"
arch=('x86_64')
url="https://github.com/alexcoder04/kherson.git"
license=('GPL3')
groups=()
depends=()
makedepends=(git go)
checkdepends=()
optdepends=()
provides=()
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=("git+$url")
noextract=()
md5sums=('SKIP')
validpgpkeys=()

package() {
  cd "$_pkgname"
  NAME="$pkgname" DESTDIR="$pkgdir/" make install
}
