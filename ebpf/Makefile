clean:
	-$(RM) dns
dns:
	clang-10 -O2 -g -Wall -Werror -emit-llvm -c dns.c -o - | llc-10 -march=bpf -mcpu=probe -filetype=obj -o dns
