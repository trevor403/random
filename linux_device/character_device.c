#define FUSE_USE_VERSION 31
#include <cuse_lowlevel.h>
#include <fuse_opt.h>

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "ioctl.h"
#include "linear.h"

static void *global_buf;
static size_t buffer_size;

static void *temp_buf;

static void srandom_open(fuse_req_t req, struct fuse_file_info *fi) {
    fuse_reply_open(req, fi);
}

static void srandom_read(fuse_req_t req, size_t size, off_t off, struct fuse_file_info *fi) {
    // printf("got a request to read buffer size %lu at offset %lu\n", size, off);
    // max read is 2**14
    if (size > 1 << 14) {
        size = 1 << 14;
    }

    // optimize large reads
     if (size > 1 << 10) {
        int rands = size / 8;
        int temp_buf_size = rands * 8;
        if (temp_buf_size != sizeof(temp_buf))
            temp_buf = realloc(temp_buf, temp_buf_size);

        // fill the buffer with random values
        GoSlice buf = {temp_buf, temp_buf_size, temp_buf_size};
        GetRandomBytesFromGo(buf);

        fuse_reply_buf(req, temp_buf, temp_buf_size);
     } else {
        // fix read size to size of buffer
        if (size > buffer_size - off)
            size = buffer_size - off;

        // get new random value
        uint64_t value = GetRandomFromGo();
        memcpy(global_buf, &value, sizeof(value));
        fuse_reply_buf(req, global_buf, size);
     }

    //  printf("sending back %lu\n", size);
}

static const struct cuse_lowlevel_ops srandom_funcs = {
    .open = srandom_open,
    .read = srandom_read,
};

int main(int argc, char **argv) {
    printf("Emulated Character Device Running...\n");

    char *foreground[] = {argv[0], "-f", NULL};
    struct fuse_args args = {2, foreground, 0};
    fuse_opt_parse(&args, NULL, NULL, NULL);

    // buffer init
    global_buf = malloc(8);
    memset(global_buf, 0xF5, buffer_size = 8); // zero out buffer

    // device name
    char dev_name[128] = "DEVNAME=";
    strncat(dev_name, "srandom", sizeof(dev_name) - strlen(dev_name));
    const char *dev_info_argv[] = {dev_name};

    // cuse options
    struct cuse_info ci;
    memset(&ci, 0, sizeof(ci));
    ci.dev_info_argc = 1;
    ci.dev_info_argv = dev_info_argv;
    ci.flags = CUSE_UNRESTRICTED_IOCTL;
    return cuse_lowlevel_main(args.argc, args.argv, &ci, &srandom_funcs, NULL);
}