output "lb_dns_name" {
  value = "${aws_lb.ecs_alb.dns_name}"
  description = "DNS name of ALB."
}