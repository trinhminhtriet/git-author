require 'fileutils'
require 'rake/clean'

PROGNAME = 'git-author'
SUPPORTED = [
  ['darwin', 'arm64'],
  ['darwin', 'amd64'],
  ['linux', 'amd64'],
  ['linux', 'arm64'],
  ['linux', 'arm'],
]
OUTDIR = 'out'
RELEASE_DIRS = SUPPORTED.map do |os, arch|
    "#{OUTDIR}/#{os}_#{arch}"
end

task default: [:build]

desc 'Run all unit tests'
task :test do
  sh 'go test ./internal/...'
end

desc 'Run go fmt'
task :fmt do
  sh 'go fmt ./internal/...'
  sh 'go fmt *.go'
end

desc 'Build executable'
task :build do
  gohostos = `go env GOHOSTOS`.strip
  gohostarch = `go env GOHOSTARCH`.strip
  build_for_platform gohostos, gohostarch
end

namespace 'release' do
  directory OUTDIR

  RELEASE_DIRS.each do |dir|
    directory dir
  end

  desc 'Build binaries for all supported platforms'
  task build: RELEASE_DIRS do
    SUPPORTED.each do |os, arch|
      output_dir = "#{OUTDIR}/#{os}_#{arch}"
      build_for_platform(os, arch, out: "#{output_dir}/#{PROGNAME}")

      version = get_version
      sh "tar czf #{OUTDIR}/gitwho_#{version}_#{os}_#{arch}.tar.gz "\
        "-C #{OUTDIR} #{os}_#{arch}"
    end
  end

  desc 'Sign checksum of built artifacts'
  task :sign do
    FileUtils.cd(OUTDIR) do
      version = get_version
      sumsfile = "SHA2-256SUMS_#{version}.txt"
      sh "shasum -a 256 **/git-author > #{sumsfile}"
      sh "ssh-keygen -Y sign -n file -f ~/.ssh/gitwho_ed25519 #{sumsfile}"
    end
  end

  task all: [:build, :sign]
end

CLOBBER.include(OUTDIR)
CLOBBER.include(PROGNAME)

def get_version()
  `git describe --tags --always --dirty`.strip
end

def get_commit()
  `git rev-parse --short HEAD`.strip
end

def build_for_platform(goos, goarch, out: PROGNAME)
  version = get_version
  rev = get_commit
  sh "GOOS=#{goos} GOARCH=#{goarch} go build -a -o #{out} "\
    "-ldflags '-s -w -X main.Commit=#{rev} -X main.Version=#{version}'"
end
