FROM temp:latest as base_img

# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM alpine:latest
WORKDIR /work
COPY --from=base_img /work/tracker /work/tracker
EXPOSE 1337
ENTRYPOINT /work/tracker