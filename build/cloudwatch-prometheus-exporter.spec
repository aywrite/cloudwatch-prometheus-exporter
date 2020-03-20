# SPEC file overview:
# https://docs.fedoraproject.org/en-US/quick-docs/creating-rpm-packages/#con_rpm-spec-file-overview
# Fedora packaging guidelines:
# https://docs.fedoraproject.org/en-US/packaging-guidelines/


Name: cloudwatch-prometheus-exporter
Version: 0.0.5
Release: 0%{?dist}
Summary: Cloudwatch Prometheus Exporter
License: BSD
URL: https://www.covergenius.com
BuildRequires: golang >= 1.12
BuildArch: x86_64


%description
Cloudwatch Prometheus Exporter


%prep
mkdir -p %{_topdir}/{BUILD,RPMS,SOURCES,SPECS,SRPMS}
cp -rf %{_sourcedir}/* %{_topdir}/BUILD


%build
go build


%install
rm -rf %{buildroot}
%{__install} -D -m 0644 %{_topdir}/BUILD/cloudwatch-prometheus-exporter %{buildroot}/%{_sbindir}/cloudwatch-prometheus-exporter


%files
%defattr(755,root,root,755)
%{_sbindir}/cloudwatch-prometheus-exporter


%changelog
* Tue Feb 11 2020 Serghei Anicheev <serghei@covergenius.com>
- Now can specify length in config.yaml
* Tue Dec 17 2019 Serghei Anicheev <serghei@covergenius.com>
- Several fixes: map protection from concurrent access and DescribeTags limit
* Tue Nov 19 2019 Serghei Anicheev <serghei@covergenius.com>
- Initial commit