
SHELL = /bin/sh
PREFIX ?= /usr
NAME = i3gocks

build:
	GOOS=linux GOARCH=amd64 go build -o $(NAME) .

clean:
	rm -f $(NAME)

install: build
	install -Dm755 "$(NAME)" "$(DESTDIR)$(PREFIX)/bin/$(NAME)"
	install -Dm644 "README.md" "$(DESTDIR)$(PREFIX)/share/doc/$(NAME)/README.md"
	install -Dm644 "LICENSE" "$(DESTDIR)$(PREFIX)/share/licenses/$(NAME)/LICENSE"
	install -Dm644 "$(NAME).1" "$(DESTDIR)$(PREFIX)/share/man/man1/$(NAME).1"

uninstall:
	$(RM) "$(DESTDIR)$(PREFIX)/bin/$(NAME)"
	$(RM) -r "$(DESTDIR)$(PREFIX)/share/doc/$(NAME)"
	$(RM) -r "$(DESTDIR)$(PREFIX)/share/licenses/$(NAME)"
	$(RM) "$(DESTDIR)$(PREFIX)/share/man/man1/$(NAME).1"

