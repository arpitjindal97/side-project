FROM temp:latest as base_img

# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM alpine:latest
WORKDIR /work
COPY --from=base_img /work/refresher /work/refresher
EXPOSE 8081
ENTRYPOINT /work/refresher